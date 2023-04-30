USE devbook;

INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Josnei Silveira", "josnei123", "josnei@hotmail.com", "$2a$10$1rCXSl4QFbM5Wbec..OzdusCFl7U2sT67K88VrwRNknsNVPJb3TuK"),
("Carlos Pafúncio", "pafascarlos", "carlospafu@bol.com", "$2a$10$1rCXSl4QFbM5Wbec..OzdusCFl7U2sT67K88VrwRNknsNVPJb3TuK"),
("Wilson Júnior", "wilsaogod", "whjcbr123@hotmail.com",  "$2a$10$1rCXSl4QFbM5Wbec..OzdusCFl7U2sT67K88VrwRNknsNVPJb3TuK");

INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);