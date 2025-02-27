export class ChatMessage {
    public userId: string;
    public message: string;
    public timestamp: number;
    public messageId: string;

    constructor(
        userId: string,
        message: string,
        timestamp: number,
        messageId: string
    ) {
        this.userId = userId;
        this.message = message;
        this.timestamp = timestamp;
        this.messageId = messageId;
    }

    static fromJSON(json: any): ChatMessage {
        return new ChatMessage(
            json.userId,
            json.message,
            json.timestamp,
            json.messageId
        );
    }
}

export class SentimentAnalysis {
    public compound: number;
    public negative: number;
    public neutral: number;
    public positive: number;

    constructor(
        compound: number,
        negative: number,
        neutral: number,
        positive: number
    ) {
        this.compound = compound;
        this.negative = negative;
        this.neutral = neutral;
        this.positive = positive;
    }

    static fromJSON(json: any): SentimentAnalysis {
        return new SentimentAnalysis(
            json.compound,
            json.negative,
            json.neutral,
            json.positive
        );
    }
}

export class ChatSentiment {
    public sentiment: SentimentAnalysis;
    public messageId: string;

    constructor(sentiment: SentimentAnalysis, messageId: string) {
        this.sentiment = sentiment;
        this.messageId = messageId;
    }

    static fromJSON(json: any): ChatSentiment {
        return new ChatSentiment(
            SentimentAnalysis.fromJSON(json.sentiment),
            json.messageId
        );
    }
}
