from fastapi import FastAPI
from pydantic import BaseModel
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity

app = FastAPI()

class MatchRequest(BaseModel):
    resume_text: str
    job_description: str

@app.post("/analyze")
def analyze(data: MatchRequest):
    vectorizer = TfidfVectorizer(stop_words="english")
    vectors = vectorizer.fit_transform(
        [data.resume_text, data.job_description]
    )
    score = cosine_similarity(vectors[0:1], vectors[1:2])[0][0]

    return {"match_percentage": round(score * 100, 2)}
