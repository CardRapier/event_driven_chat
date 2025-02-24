from nltk.sentiment import SentimentIntensityAnalyzer #type: ignore


class Analysis():
    def __init__(self):
        self.sia = SentimentIntensityAnalyzer()

    def analyze(self, message):
        return self.sia.polarity_scores(message)