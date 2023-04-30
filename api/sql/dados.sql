USE devbook;

INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Josnei Silveira", "josnei123", "josnei@hotmail.com", "2a$10$10qb0hN3ybIhnDOeViVKQeeU7TRW1N2Zu5bpoy4sDZy9o6UmzumBi"),
("Carlos Pafúncio", "pafascarlos", "carlospafu@bol.com", "2a$10$10qb0hN3ybIhnDOeViVKQeeU7TRW1N2Zu5bpoy4sDZy9o6UmzumBi"),
("Wilson Júnior", "wilsaogod", "whjcbr123@hotmail.com",  "2a$10$10qb0hN3ybIhnDOeViVKQeeU7TRW1N2Zu5bpoy4sDZy9o6UmzumBi");

INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);