package main

import (
	"errors"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Example: "init risk_address1",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "risk_address1":
			_, err := conn.Exec(cmd.Context(), `
				DROP TABLE IF EXISTS risk_address1;
				CREATE TABLE risk_address1
				(
						address VARCHAR          NOT NULL,
						version BIGINT DEFAULT 1 NOT NULL,
						PRIMARY KEY (address, version)
				);
			`)
			return err
		case "risk_address2":
			_, err := conn.Exec(cmd.Context(), `
				DROP TABLE IF EXISTS risk_address2;
				CREATE TABLE risk_address2
				(
						id      BIGSERIAL CONSTRAINT risk_address_pk PRIMARY KEY,
						address VARCHAR          NOT NULL,
						version BIGINT DEFAULT 1 NOT NULL
				);
				
				CREATE UNIQUE INDEX risk_address2_address_version_uindex
						ON risk_address2 (address ASC, version DESC);
			`)
			return err
		default:
			return errors.New("invalid table")
		}
	},
}
