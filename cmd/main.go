package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var conn *pgx.Conn

var rootCmd = &cobra.Command{
	Use: "app",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		return err
	},
}

func main() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(insertCmd)
	rootCmd.AddCommand(selectCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
