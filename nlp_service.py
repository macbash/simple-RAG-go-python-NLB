from flask import Flask, request, jsonify
from textblob import TextBlob

app = Flask(__name__)

@app.route('/analyze_sentiment', methods=['POST'])
def analyze_sentiment():
    data = request.get_json(force=True)
    text = data['text']
    blob = TextBlob(text)
    sentiment_polarity = blob.sentiment.polarity
    sentiment_subjectivity = blob.sentiment.subjectivity
    return jsonify({
        "polarity": sentiment_polarity,
        "subjectivity": sentiment_subjectivity
    })

if __name__ == '__main__':
    app.run(port=5000)