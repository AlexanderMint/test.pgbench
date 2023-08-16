package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:     "select",
	Example: "select risk_address1",
	RunE: func(cmd *cobra.Command, args []string) error {
		table := args[0]
		if table == "" {
			return errors.New("invalid table")
		}

		var num int64
		row := conn.QueryRow(cmd.Context(), fmt.Sprintf("SELECT COUNT(*) FROM %s", table))
		err := row.Scan(&num)
		if err != nil {
			return err
		}

		if num == 0 {
			return errors.New("num is zero")
		}
		query := fmt.Sprintf("SELECT 1 FROM %s WHERE address = $1 AND version = $2 LIMIT 1", table)

		uuids := make([]string, num)
		for i := range uuids {
			id, err := uuid.Parse(fmt.Sprintf("00000000-0000-0000-0000-%012d", i))
			if err != nil {
				return err
			}
			uuids[i] = id.String()
		}

		log.Info().Int64("num", num).Send()
		start := time.Now()
		part := time.Now()
		partNum := int64(0)

		for i := range uuids {
			var b int
			r := conn.QueryRow(cmd.Context(), query, uuids[i], i)
			err := r.Scan(&b)
			if err != nil || b == 0 {
				panic(err)
			}

			if i%10_000 == 0 && i > 0 {
				log.Info().
					Str("duration", time.Since(start).String()).
					Str("average", time.Duration(time.Since(start).Nanoseconds()/int64(i)).String()).
					Str("average part", time.Duration(time.Since(part).Nanoseconds()/(int64(i)-partNum)).String()).
					Int64("i", int64(i)).
					Send()
				part = time.Now()
				partNum = int64(i)
			}
		}

		return nil
	},
}
