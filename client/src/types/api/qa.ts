//APIの仕様でバックエンドと合わせる

export type QA = {
    qid: number,
    aid: number | null,
    question_content: string,
    answer_content: string,
    // questions: Question[],
    // answers:Answer[],
};

