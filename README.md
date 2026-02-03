"# AI-Resume-Screening-System" 
curl -X POST "http://localhost:8181/match" ^
-H "Content-Type: application/json" ^
-d "{\"resume_text\":\"Python ML Docker\",\"job_description\":\"Python ML\"}"
