CREATE TABLE candidatos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    habilidades TEXT[] NOT NULL,
    status VARCHAR(50) NOT NULL
);
