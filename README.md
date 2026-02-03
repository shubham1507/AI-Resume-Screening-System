"# AI-Resume-Screening-System" 
cd ai-service
uvicorn app:app --reload

cd backend
go run main.go

cd frontend
python -m http.server 5500

