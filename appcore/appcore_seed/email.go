package appcore_seed

import (
	"case-management/model"
	"strings"

	"gorm.io/gorm"
)

func findAndExtendLongestTemplate(templates []model.EmailTemplate) (string, string) {
	var longest string
	var longestName string

	for _, template := range templates {
		if len(template.Template) > len(longest) {
			longest = template.Template
			longestName = template.Template
		}
	}

	// เพิ่ม 300 ตัวอักษรเข้าไป
	longest += strings.Repeat(" ", 300)

	return longestName, longest
}

func SeedEmailTemplates(db *gorm.DB) {
	templates := []model.EmailTemplate{
		{
			Template: "Change Passport Infomation",
			Subject:  "Change Passport Infomation",
			Body: `Deer <Name>, 

			We are in the process of updating your records in our Case Management System. Kindly provide a copy of your new passport for our records.
			If you have any questions, feel free to contact us.

			Best regrads,
			AEON
			`,
		},
		{
			Template: "Forget Password",
			Subject:  "Reset Your Password",
			Body: `Dear <Name>,
					You have requested to reset your password. Please click on the link below to set a new password:

					<ResetLink>

					This link is valid for 15 minutes. Please do not share this link with anyone for your security. If you did not request a password reset or suspect
					any unauthorized activity, please contact our support team immediately.

					Best regards,
					Your Team
			`,
		},
	}

	// ค้นหา template ที่ยาวที่สุด
	longestName, extended := findAndExtendLongestTemplate(templates)

	// อัปเดต template ที่ยาวที่สุด
	for i := range templates {
		if templates[i].Template == longestName {
			templates[i].Template = extended
			break
		}
	}

	// Insert หรือ Update
	for _, template := range templates {
		db.FirstOrCreate(&template, model.EmailTemplate{Template: template.Template})
	}
}
