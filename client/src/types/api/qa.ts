//APIの仕様でバックエンドと合わせる

export type QA = {
    qid: number,
    aid: number | null,
    question_content: string,
    answer_content: string,
    // questions: Question[],
    // answers:Answer[],
};

export type Answer = {
    id: number | null;
    aid: number,
    answer_content: string,
}

export type Question = {
    qid: number,
    question_content: string,
}