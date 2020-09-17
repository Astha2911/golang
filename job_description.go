package fixes

import (
	"bytes"
	"fmt"
	"html"
	"io"
//	"os"
	"log"
	"strings"
	"tutree/jobportal/config"
	"tutree/jobportal/models"
	"tutree/jobportal/utility"

	hparse "golang.org/x/net/html"
)

func getJobs(o int, jobID string) []models.Job {
	db := config.GetDB()

	offset := fmt.Sprintf("%d", o)

	where := "1=1"
	if jobID != "all" {
		where = "id = " + jobID
	}

	query := `
		SELECT 
			id, 
			html_job_description
		FROM 
			jobs.job
		WHERE ` +
			where + `
		OFFSET ` + offset + `
		LIMIT 50`

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Query failed.....")
		log.Println(err.Error())
		return nil
  }
  jobs := []models.Job{}
  for rows.Next() {
    var jID int
		var htmlJobDesc string

    rows.Scan(&jID, &htmlJobDesc)

    jobs = append(jobs, models.Job{ID: jID, HTMLJobDesc: htmlJobDesc})
  }
	// log.Printf("Got jobs %d", len(jobs))
  return jobs
}

func updateJobDescription(jobs []models.Job) {
	db := config.GetDB()

	for _, job := range jobs {
		query := `
			UPDATE 
				jobs.job 
			SET 
				text_job_description = $1,
				html_job_description = $2  
			WHERE 
				id = $3`

		_, err := db.Exec(query, job.TextJobDesc, job.HTMLJobDesc, job.ID)
		if err != nil {
			log.Printf("Update query failed %+v", err)
		}
	}
}

func removeAnchorTag(htmlDesc string) string {
	doc, err := hparse.Parse(strings.NewReader(htmlDesc))
	if err != nil {
			log.Println(err)
			return htmlDesc
	}
	var f func(*hparse.Node)
	f = func(n *hparse.Node) {
		if n.Type == hparse.ElementNode && n.Data == "a" {
			// Do something with n...
			for i, k := range n.Attr {
				if k.Key == "href" {   //replacing href value with #
					n.Attr[i].Val = "#"
				}
			}
			
			n.Data = "span" //coverting anchor tag to span

		}
			
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	
	var buf bytes.Buffer
  w := io.Writer(&buf)
  hparse.Render(w, doc)
  return buf.String()
}
func removeStyleAttr(htmlDesc string) string {

	doc, err := hparse.Parse(strings.NewReader(htmlDesc))
	if err != nil {
			log.Println(err)
			return htmlDesc
	
	}
	//doc, err := hparse.Parse(strings.NewReader(htmlDesc))
      //if err != nil {
		//	  panic(err)
		//	  return htmlDesc
		//	  RemoveStyleAttr(htmlDesc)
	      // hparse.Render(os.Stdout, doc)
      //}

	var f func(*hparse.Node)
	f = func(n *hparse.Node) {
	i := -1
	for index, attr := range n.Attr {
			if attr.Key == "style" {
					i = index
					break
			}
	}
	if i != -1 {
			n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
	}
	
	}
	f(doc)
	
	var buf bytes.Buffer
  w := io.Writer(&buf)
  hparse.Render(w, doc)
  return buf.String()
}


func fixHTMLDescription(htmlDesc string) string {
	hjd := html.UnescapeString(htmlDesc)	
	hjd = removeAnchorTag(hjd)
	hjd = removeStyleAttr(hjd)
	return hjd
}

//FixJobDescription fix  description for jobs
func FixJobDescription(jobID string) []int {
	jobsOffset := 0
	jobs := getJobs(jobsOffset, jobID)
	
	fixedJobs := []int{}
	
	for len(jobs) > 0 {
		for i := range jobs {
			hd := jobs[i].HTMLJobDesc
			hjd := fixHTMLDescription(hd)
			tjd := utility.HTML2Text(hjd)

			jobs[i].HTMLJobDesc = hjd
			jobs[i].TextJobDesc = tjd
			fixedJobs = append(fixedJobs, jobs[i].ID)

		}

		updateJobDescription(jobs)
		jobsOffset += len(jobs)
		jobs = getJobs(jobsOffset, jobID)
	}

	return fixedJobs
}