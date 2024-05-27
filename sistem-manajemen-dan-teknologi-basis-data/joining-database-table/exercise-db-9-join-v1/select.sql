-- TODO: answer here
SELECT
    r.id AS id,
    s.fullname AS fullname,
    s.class AS class,
    s.status AS status,
    r.study AS study,
    r.score AS score
FROM reports r
JOIN students s ON s.id = r.student_id
WHERE r.score < 70
ORDER BY r.score ASC; 