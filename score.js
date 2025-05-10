const submitScore(name, score, time) {
    fetch("http://localhost:8080/score", {
        method: "POST",
        headers: {
            "Content-type": "application/json"
        },
        body: JSON.stringify({ name, score, time })
    })
        .then(res => res.json())
        .then(data => {
            console.log("Score Submited:", data)
            fetchScores()
        })
        .catch(err => console.log("Error in submiting", err))
}

const fetchScores() {
    fetch("http://localhost:8080/scores")
        .then(res => res.json())
        .then(data => {
            console.log("Fetched scores:", data);
            displayScores(data); // TODO: write this function
        })
        .catch(err => console.error("Error fetching scores:", err));
}
