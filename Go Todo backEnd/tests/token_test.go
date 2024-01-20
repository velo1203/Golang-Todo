package octet_test

// type TokenTestSuite struct {
// 	suite.Suite
// 	db         *gorm.DB
// 	repository repositories.TokenRepository
// 	service    services.TokenService
// 	controller controllers.TokenController
// }

// func (suite *TokenTestSuite) SetupSuite() {
// 	err := os.Remove("test.db")
// 	if err != nil {
// 		suite.Fail("Failed to remove the test database file")
// 	}

// 	database.Init()
// 	gorm_config := gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}}
// 	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm_config)
// 	repository := repositories.NewTokenRepository(db)
// 	service := services.NewTokenService(repository)
// 	controller := controllers.NewTokenController(service)

// 	suite.db = db
// 	suite.service = service
// 	suite.repository = repository
// 	suite.controller = controller

// 	suite.CreateSampleData()
// }

// func (suite *TokenTestSuite) CreateSampleData() {
// 	suite.db.AutoMigrate(&models.Token{})
// 	var tokens []models.Token
// 	for i := 1; i < 100; i++ {
// 		expire_at := time.Now().Add(time.Hour * 24 * time.Duration(i+45))
// 		tokenString, _ := suite.service.Generate("superrichquiz_octet", expire_at.Unix())
// 		token := models.Token{
// 			ID:        uuid.NewString(),
// 			Token:     tokenString,
// 			Status:    "valid",
// 			ExpireAt:  expire_at,
// 			UpdatedAt: time.Now(),
// 			CreatedAt: time.Now(),
// 		}
// 		tokens = append(tokens, token)
// 	}

// 	for _, token := range tokens {
// 		suite.db.Create(&token)
// 	}
// }

// func (suite *TokenTestSuite) TearDownSuite() {
// 	// suite.db.Migrator().DropTable(&models.Token{})

// 	// err := os.Remove("test.db")
// 	// if err != nil {
// 	// 	suite.Fail("Failed to remove the test database file")
// 	// }
// }
// func (suite *TokenTestSuite) SetupTest() {
// 	// suite.userRepository.DeleteByName("testUserName0001")
// }

// func (suite *TokenTestSuite) TearDownTest() {

// }

// func TestTokenSuite(t *testing.T) {
// 	suite.Run(t, new(TokenTestSuite))
// }

// // 토큰 생성 테스트
// func (suite *TokenTestSuite) TestGenerateTokenSuccess() {
// 	user_id := "testuser"
// 	expire_at := time.Now().Add(time.Hour * 24 * 365).Unix()
// 	token, err := suite.service.Generate(user_id, expire_at)
// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), token)
// }

// // 토큰 생성 테스트
// func (suite *TokenTestSuite) TestGenerateTokenString() {
// 	user_id := "superrichquiz_octet"
// 	expire_at := time.Now().Add(time.Hour * 24 * 365 * 100).Unix()
// 	tokenString, err := suite.service.Generate(user_id, expire_at)
// 	assert.NoError(suite.T(), err)

// 	jwtToken, err := suite.service.VerifyTokenString(tokenString)
// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), jwtToken)
// }

// // 토큰 검증 테스트
// func (suite *TokenTestSuite) TestVerifyTokenString() {
// 	user_id := "testuser"
// 	expire_at := time.Now().Add(time.Hour * 24 * 365).Unix()
// 	tokenString, err := suite.service.Generate(user_id, expire_at)
// 	assert.NoError(suite.T(), err)

// 	jwtToken, err := suite.service.VerifyTokenString(tokenString)
// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), jwtToken)
// 	assert.Equal(suite.T(), tokenString, jwtToken.Raw)
// }

// // 토큰 유효성 테스트
// func (suite *TokenTestSuite) TestValidToken() {
// 	user_id := "testuser"
// 	expire_at := time.Now().Add(time.Hour * 24 * 365).Unix()
// 	tokenString, err := suite.service.Generate(user_id, expire_at)
// 	assert.NoError(suite.T(), err)

// 	jwtToken, err := suite.service.VerifyTokenString(tokenString)
// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), jwtToken)

// 	valid, err := suite.service.ValidToken(jwtToken)
// 	assert.NoError(suite.T(), err)
// 	assert.True(suite.T(), valid)
// }

// // 토큰 유효성 테스트 : Expire
// func (suite *TokenTestSuite) TestExpiredToken() {
// 	user_id := "testuser"
// 	expire_at := time.Now().Add(time.Hour * 1).Unix()
// 	tokenString, err := suite.service.Generate(user_id, expire_at)
// 	assert.NoError(suite.T(), err)

// 	jwtToken, err := suite.service.GetJwtToken(tokenString)
// 	assert.NoError(suite.T(), err)
// 	// assert.Nil(suite.T(), jwtToken)

// 	expirationTime, err := jwtToken.Claims.GetExpirationTime()
// 	assert.NoError(suite.T(), err)
// 	assert.Greater(suite.T(), expirationTime.Time, time.Now())

// 	expire_at = time.Now().Add(-time.Hour * 1).Unix()
// 	tokenString, err = suite.service.Generate(user_id, expire_at)
// 	assert.NoError(suite.T(), err)

// 	jwtToken, err = suite.service.GetJwtToken(tokenString)
// 	assert.Error(suite.T(), err)
// 	// assert.Nil(suite.T(), jwtToken)

// 	expirationTime, err = jwtToken.Claims.GetExpirationTime()
// 	assert.NoError(suite.T(), err)
// 	assert.Greater(suite.T(), time.Now(), expirationTime.Time)
// }

// func (suite *TokenTestSuite) TestCreateToken() {
// 	dateString := "2023-08-15T01:50:08.706Z"
// 	layout := "2006-01-02T15:04:05.000Z"
// 	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklkeCI6MTc2NTQyLCJ0b2tlblR5cGUiOiJXQUxMRVQiLCJ0b2tlbkV4cGlyZWREYXRlIjoiMjAyMy0wOC0xNVQwMTo1MDowOC43MDZaIn0.skgP6ysLNx6KRDBYZy3miZW1Q95iV_Cw0ZVWXj1tJyw"
// 	expire_at, _ := time.Parse(layout, dateString)
// 	token := models.Token{
// 		ID:       uuid.NewString(),
// 		Token:    tokenString,
// 		Status:   "valid",
// 		ExpireAt: expire_at,
// 	}
// 	result, err := suite.repository.Create(&token)

// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), tokenString, result.Token)
// }

// // 리프레시할 토큰 목록
// func (suite *TokenTestSuite) TestListOfTokensToRefresh() {
// 	tokens, err := suite.repository.ListRefresh()
// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), tokens)
// 	assert.Greater(suite.T(), len(*tokens), 0)
// }

// // 토큰 리프레시
// func (suite *TokenTestSuite) TestOldTokenRefresh() {
// 	tokens, err := suite.repository.ListRefresh()
// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), tokens)

// 	for _, token := range *tokens {
// 		token.Status = "expired"
// 		result, err := suite.repository.Update(&token)
// 		assert.NoError(suite.T(), err)
// 		assert.Equal(suite.T(), "expired", result.Status)

// 		user_id := "superrichquiz_octet"
// 		expire_at := result.ExpireAt.Add(time.Hour * 24 * 90)
// 		tokenString, err := suite.service.Generate(user_id, expire_at.Unix())
// 		assert.NoError(suite.T(), err)

// 		refreshedToken := models.Token{
// 			ID:       uuid.NewString(),
// 			Token:    tokenString,
// 			Status:   "valid",
// 			ExpireAt: expire_at,
// 		}

// 		newToken, err := suite.repository.Create(&refreshedToken)
// 		assert.NoError(suite.T(), err)
// 		assert.Equal(suite.T(), tokenString, newToken.Token)
// 	}
// }

// func (suite *TokenTestSuite) TestResponseToken() {
// 	respBody := []byte(`{
//     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklkeCI6MTc1MzU4LCJ0b2tlblR5cGUiOiJXQUxMRVQiLCJ0b2tlbkV4cGlyZWREYXRlIjoiMjAyMy0wOC0xNFQwNToxMTo1Ni41MzhaIn0.e1i4_y8ItC8Vje13Ew8NHwZTElOMBObIZjpGLgRCdyE",
//     "tokenBody": {
//         "tokenExpiredDate": "2023-08-13T06:06:12.065Z",
//         "tokenIdx": 174025,
//         "tokenType": "WALLET"
//     }
// 	}`)

// 	var response map[string]interface{}
// 	err := json.Unmarshal(respBody, &response)
// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), response)
// 	assert.Equal(suite.T(), response["token"], "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklkeCI6MTc1MzU4LCJ0b2tlblR5cGUiOiJXQUxMRVQiLCJ0b2tlbkV4cGlyZWREYXRlIjoiMjAyMy0wOC0xNFQwNToxMTo1Ni41MzhaIn0.e1i4_y8ItC8Vje13Ew8NHwZTElOMBObIZjpGLgRCdyE")

// 	tokenBody := response["tokenBody"].(map[string]interface{})
// 	tokenExpiredDate := tokenBody["tokenExpiredDate"].(string)
// 	assert.Equal(suite.T(), "2023-08-13T06:06:12.065Z", tokenExpiredDate)
// }

// func (suite *TokenTestSuite) TestResponseError() {
// 	respBody := []byte(`{
//     "errorCode": "ERR_0105001",
//     "message": "Invalid token"
// 	}`)

// 	var response map[string]interface{}
// 	err := json.Unmarshal(respBody, &response)
// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), response)

// 	assert.NotNil(suite.T(), response["errorCode"])
// 	assert.Equal(suite.T(), "ERR_0105001", response["errorCode"].(string))

// 	if response["errorCode"] != nil && response["errorCode"].(string) == "ERR_0105001" {
// 		assert.True(suite.T(), true)
// 	}
// }
