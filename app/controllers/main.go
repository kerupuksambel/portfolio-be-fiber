package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	Model "github.com/kerupuksambel/portfolio-be/app/models"
)

func WorkExperience(c *fiber.Ctx) error {
	// Static data here...
	var workExperiences []Model.WorkExperience = []Model.WorkExperience{
		{
			Name:     "Rally the Locals",
			Position: "Lead Developer",
			Location: "Winnipeg, Canada",
			Date: Model.Date{
				Start: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				End:   nil,
			},
			Achievements: []string{
				"Developed Rally the Locals main frontend along with supporting pages",
				"Developed the payment gateway using Stripe",
				"Developed and improved RTL admin dashboard for better administrator and easier organization on the admin side",
				"Designed and led the migration from PHP-based backend with Laravel to NodeJS + ExpressJS backend",
				"Designed the pipeline in Bitbucket to implement Continuous Integration in development side",
				"Handled the deployment for our migrated frontend and backend",
			},
			TechStacks: []Model.TechStack{
				{
					Name: "Laravel",
					Icon: "devicon:laravel",
					Link: "https://laravel.com/",
				},
				{
					Name: "NextJS",
					Icon: "devicon:nextjs",
					Link: "https://nextjs.org/",
				},
				{
					Name: "ExpressJS",
					Icon: "skill-icons:expressjs-light",
					Link: "https://expressjs.com/",
				}
				{
					Name: "Tailwind",
					Icon: "devicon:tailwindcss",
					Link: "https://tailwindcss.com",
				},
				{
					Name: "Stripe",
					Icon: "logos:stripe",
					Link: "https://stripe.com/",
				},
				{
					Name: "Docker",
					Icon: "logos:docker-icon",
					Link: "https://docker.com/",
				},
				{
					Name: "Atlassian Suite",
					Icon: "logos:atlassian",
					Link: "https://atlassian.com/",
				},
			},
		},
		{
			Name:     "Technowand",
			Position: "Freelance Full Stack Web Developer",
			Location: "Australia",
			Date: Model.Date{
				Start: time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			},
			Achievements: []string{
				"Developed fitness and personal trainer website with CodeIgniter",
				"Handled the Stripe integration on the site for easier transaction handling",
				"Developed the application using Scrum methodology ",
			},
			TechStacks: []Model.TechStack{
				{
					Name: "CodeIgniter",
					Icon: "logos:codeigniter-icon",
					Link: "https://codeigniter.com",
				},
				{
					Name: "Stripe",
					Icon: "logos:stripe",
					Link: "https://stripe.com/",
				},
				{
					Name: "Bootstrap",
					Icon: "devicon:bootstrap",
					Link: "http://getbootstrap.com/",
				},
				{
					Name: "Jira",
					Icon: "logos:jira",
					Link: "https://www.atlassian.com/software/jira",
				},
			},
		},
		{
			Name:     "DPTSI - Institut Teknologi Sepuluh Nopember",
			Position: "Intern Full Stack Web Developer",
			Location: "Surabaya, Indonesia",
			Date: Model.Date{
				Start: time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC),
				End:   time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC),
			},
			Achievements: []string{
				"Developed the proctoring system for ITS entrance exam system",
				"Designed the database and the data management system to recapitulate and generate reports of test results faster and more efficiently, with load speed increase up to 60%.",
				"Developed a MOOC (Massive Open Online Course) for ITS using Laravel and Bootstrap",
				"Designed the database and the models to accommodate the scalable MOOC application",
			},
			TechStacks: []Model.TechStack{
				{
					Name: "Laravel",
					Icon: "devicon:laravel",
					Link: "https://laravel.com/",
				},
				{
					Name: "Microsoft SQL Server",
					Icon: "devicon:microsoftsqlserver",
					Link: "https://www.microsoft.com/en-us/sql-server/",
				},
				{
					Name: "Bootstrap",
					Icon: "devicon:bootstrap",
					Link: "http://getbootstrap.com/",
				},
			},
		},
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    workExperiences,
	})
}
