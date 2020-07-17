package cmd

import (
	"fmt"

	"github.com/imrancluster/mama/config"
	"github.com/imrancluster/mama/conn"
	"github.com/spf13/cobra"
)

var msisdn string

// deleteEnrolmentCmd represents the all command
var deleteEnrolmentCmd = &cobra.Command{
	Use:   "enrolment",
	Short: "Enrolment short",
	Long:  `Enrolment long`,
	Run: func(cmd *cobra.Command, args []string) {
		mobile, _ := cmd.Flags().GetString("msisdn")

		fmt.Printf("\nDelete Operation is running for msisdn: %s\n\n", mobile)

		// Deleting enrolment records
		RunDeleteOperation(mobile)

		// Operation END message
		fmt.Printf("\n\n*** DONE ***\n\n")
	},
}

func init() {
	deleteCmd.AddCommand(deleteEnrolmentCmd)

	deleteEnrolmentCmd.Flags().StringVarP(&msisdn, "msisdn", "m", "", "Use valid mobile number, eg: 8801799997163")
}

// RunDeleteOperation to delete enrolment record by msisdn
func RunDeleteOperation(msisdn string) {

	// Delete record form enrolment server
	config.Init()
	conn.ConnectEnrolmentDB()

	fmt.Println("Action:")

	db := conn.EnrolmentDB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("DB Connection ERROR:", err)
	}

	sql := fmt.Sprintf(
		`
	DELETE FROM dh_corona_enrolment WHERE msisdn='%s';
	DELETE FROM dh_corona_family_member WHERE msisdn='%s';
	DELETE dh_corona_claims, dh_corona_claim_sms, dh_corona_claim_documents
		FROM dh_corona_claims
			LEFT JOIN dh_corona_claim_sms ON dh_corona_claim_sms.claim_id = dh_corona_claims.claim_id
			LEFT JOIN dh_corona_claim_documents ON dh_corona_claim_documents.claim_id = dh_corona_claims.claim_id
				WHERE dh_corona_claims.member_msisdn=SUBSTRING('%s', 3);
		`,
		msisdn, msisdn, msisdn,
	)

	fmt.Println("* Deleting records from enrolment server")
	_, exerr := db.Exec(sql)
	check(exerr)

	// Delete records from core db
	conn.ConnectCoreDB()
	coreDB := conn.CoreDB()
	defer coreDB.Close()

	err = coreDB.Ping()
	check(err)

	// selectSQL string
	selectSQL := fmt.Sprintf(`SELECT membership_no FROM msisdns WHERE msisdns.msisdn='%s' LIMIT 1;`, msisdn)
	rows, err := coreDB.Query(selectSQL)
	defer rows.Close()
	check(err)

	var membershipNo string
	for rows.Next() {
		err := rows.Scan(&membershipNo)
		check(err)
	}

	fmt.Println("* Deleting records from CORE DB")
	deleteSQL := fmt.Sprintf(
		`
		DELETE FROM "members" WHERE "membership_no" IN ('%s');
		DELETE FROM "profiles" WHERE "membership_no" IN ('%s');
		DELETE FROM "member_verification" WHERE "membership_no" IN ('%s');
		DELETE FROM "subscription_history" WHERE "membership_no" IN ('%s');
		DELETE FROM "membership_history" WHERE "membership_no" IN ('%s');
		DELETE FROM "transactions" WHERE "membership_no" IN ('%s');
		DELETE FROM "retailer_invites" WHERE "msisdn"= '%s';
		DELETE FROM "msisdns" WHERE "membership_no" IN ('%s');
		DELETE FROM "gift_avails" WHERE "msisdn" = '%s';
		`, membershipNo, membershipNo, membershipNo, membershipNo, membershipNo, membershipNo, msisdn, membershipNo, msisdn,
	)

	_, err = coreDB.Exec(deleteSQL)
	check(exerr)

	fmt.Println()
	fmt.Println("RESULTS:")
	fmt.Println("* enrolment record has been deleted")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
