CREATE TABLE UserAccessToken (
    UserID STRING(36) NOT NULL,
    Score INT64 NOT NULL,
    Token STRING(1000) NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true)
) PRIMARY KEY(UserID);
