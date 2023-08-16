package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:     "insert",
	Example: "insert risk_address1 10000",
	RunE: func(cmd *cobra.Command, args []string) error {
		table := args[0]
		if table == "" {
			return errors.New("invalid table")
		}

		max, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return err
		}

		var num int64
		row := conn.QueryRow(cmd.Context(), fmt.Sprintf("SELECT COUNT(*) FROM %s", table))
		err = row.Scan(&num)
		if err != nil {
			return err
		}

		if max < num {
			return errors.New("max < num")
		}

		query := fmt.Sprintf("INSERT INTO %s (address, version) VALUES ($1, $2)", table)

		uuids := make([]string, max)
		for i := range uuids {
			id, err := uuid.Parse(fmt.Sprintf("00000000-0000-0000-0000-%012d", i))
			if err != nil {
				return err
			}
			uuids[i] = id.String()
		}

		log.Info().Int64("num", num).Int64("max", max).Send()
		start := time.Now()
		part := time.Now()
		partNum := num

		for i := num; i < max; i++ {
			conn.Exec(cmd.Context(),
				query,
				uuids[i], i,
			)
			if i%10_000 == 0 && i != num {
				log.Info().
					Str("duration", time.Since(start).String()).
					Str("average", time.Duration(time.Since(start).Nanoseconds()/(i-num)).String()).
					Str("average part", time.Duration(time.Since(part).Nanoseconds()/(i-partNum)).String()).
					Int64("num", i+num).
					Send()
				part = time.Now()
				partNum = i
			}
		}

		return nil
	},
}
