DROP TABLE IF EXISTS tokens;

CREATE TABLE tokens (
    chain_id INT NOT NULL,
    token_id BIGINT UNSIGNED NOT NULL,
    name_hash INT NOT NULL,
    data TEXT NOT NULL,
    minter CHAR(40) NOT NULL,
    veracity INT NOT NULL, 
    PRIMARY KEY (chain_id, token_id)
);

CREATE INDEX tokens_name_hash_index ON tokens (name_hash);
CREATE INDEX tokens_minter_index ON tokens (minter);

DROP TABLE IF EXISTS lookup;

CREATE TABLE lookup (
    word_hash INT NOT NULL,
    weight INT NOT NULL,
    chain_id INT NOT NULL,
    token_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (word_hash, weight, chain_id, token_id)
);
