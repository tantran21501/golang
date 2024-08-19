
CREATE TABLE confirmed_request(
   confirmed_request_id UUID NOT NULL
  ,tenant_id UUID NOT NULL
  ,user_id UUID NOT NULL
  ,create_by UUID NOT NULL
  ,create_time TIMESTAMP(6) NOT NULL
  ,update_by UUID NOT NULL
  ,update_time TIMESTAMP(6) NOT NULL

  ,PRIMARY KEY(confirmed_request_id)
);
