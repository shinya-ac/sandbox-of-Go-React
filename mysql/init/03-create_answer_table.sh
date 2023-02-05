#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "CREATE TABLE answers (
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  question_id INT NOT NULL,
  user_id INT NOT NULL,
  answer_content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(ID),
  FOREIGN KEY (question_id) REFERENCES questions(ID),
  UNIQUE (question_id)
);"