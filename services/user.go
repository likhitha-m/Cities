package services

import (
	"context"
	"fmt"
	"os"

	"time"

	// "fmt"
	// "time"

	"cities/config"
	"cities/models"
	"cities/types"
	"cities/utils"

	// "cities/types"

	"cities/storage"

	"github.com/dgrijalva/jwt-go"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	"strconv" 
)

type UsersReceiver struct {
	// MDB        *mongo.Database
	// CustomerId int

	// UserPayload types.CityPayload
}

func GetUserByEmail(email string) (models.UserModel, error) {
	var user models.UserModel
	mdb := storage.MONGO_DB

	filter := bson.M{
		"email": email,
	}

	result := mdb.Collection(models.UsersCollection).FindOne(context.TODO(), filter)
	err := result.Decode(&user)
	if err != nil {
		logger.Error("func_S_GetGrant: Error in ", err)
		return user, err
	}
	return user, nil
}
func CreateUser(c *types.UserPayload) (interface{}, error) {
	um := models.UserModel{}
	um.Email = utils.ToLowerCase(c.Email)
	encPassword, err := utils.Encrypt(c.Password, os.Getenv("PASSWORD_ENC_KEY"))
	if err != nil {
		logger.Error("func_CreateCustomer: Error in encrypt password: ", err)
		return nil, err
	}
	um.Password = encPassword
	um.Name = c.Name
	um.MobileNumber = c.MobileNumber
	um.CreatedAt = time.Now().UTC()
	um.UpdatedAt = time.Now().UTC()
	mdb := storage.MONGO_DB
	_, err = mdb.Collection(models.UsersCollection).InsertOne(context.TODO(), um)
	if err != nil {
		logger.Error("func_AddCity: ", err)
		return nil, err
	}
	SendSignupVerificationLink(um)

	return nil, nil
}

func SendSignupVerificationLink(user models.UserModel) error {
	hashKey, err := utils.SimpleEncrypt(user.Email)
	fmt.Println("err-->>", err)

	if err != nil {
		return err
	}
	fmt.Println("hashkey-------------------------------->>>>", hashKey)
	vse := &types.VerifySignupEmailTemplate{}
	vse.VerifySignupLink = os.Getenv("BASE_PATH") + "/" + os.Getenv("VERIFY_EMAIL_PAGE") + "/" + hashKey
	// vse.PrivacyPolicyLink = os.Getenv("BASE_PATH") + "/" + os.Getenv("PRIVACY_POLICY_PAGE")
	// vse.TOSLink = os.Getenv("BASE_PATH") + "/" + os.Getenv("TERMS_OF_SERVICE_PAGE")
	// vse.UnsubscribeLink = ""
	templateData, err := utils.GetStringify(vse)
	if err != nil {
		logger.Error("func_SignupVerificationLink: Error in get stringify: ", err)
		return err
	}

	r := utils.NewSendTemplateEmailReciever(os.Getenv("EMAIL"), user.Email, config.VerifySignupEmailTemplateName, templateData)
	go r.SendTemplateEmail()

	return nil
}
func VerifyEmail(hashKey string) error {

	email, err := utils.SimpleDecrypt(hashKey)
	if err != nil {
		logger.Error("func_VerifyEmail: Error in Decrypt. Error: ", err)
		return err
	}
	user, err := GetUserByEmail(email)
	if err != nil {
		logger.Error("func_VerifyEmail: Error in GetCustomerByEmail. Error: ", err)
		return err
	}

	if user.IsVerified {
		return config.ErrEmailAlreadyVerified
	}

	user.IsVerified = true
	if err := UpdateUser(user); err != nil {
		logger.Error("func_VerifyEmail: Error in update customer. Error: ", err)
		return err
	}

	return nil
}

func UpdateUser(user models.UserModel) error {

	mdb := storage.MONGO_DB

	filter := bson.M{
		"email": user.Email,
	}
	update := bson.M{"$set": bson.M{"is_verified": user.IsVerified}}

	fmt.Println("update --->>>", update, filter)

	_, err := mdb.Collection(models.UsersCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func Login(payload types.LoginBody) (types.LoginOutput, error) {
	var loginOutput types.LoginOutput
	user, err := GetUserByEmail(payload.Email)
	if err != nil {
		logger.Error("IdentityByEmail: Error in get customer by email. Error: ", err)
		return loginOutput, err
	}
	if user.Email == "" {
		return loginOutput, config.ErrUserDoesNotExist
	}
	if !user.IsVerified {
		return loginOutput, config.ErrEmailNotVerified
	}
	token , err := GenerateToken(user)
	if err != nil {
		logger.Error("GenerateToken: Error in generating the token Error: ", err)
		return loginOutput, err
	}
	loginOutput.Email = user.Email
	loginOutput.Name= user.Name
	loginOutput.Id= user.ID
	loginOutput.Token  = token

	return loginOutput, nil
}

func  GenerateToken(userResult models.UserModel) (string, error) {
	// Generate Token
	token := jwt.New(jwt.SigningMethodHS256)
	tokenExpiredBy, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRY"))
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userResult.Name
	claims["user_id"] = userResult.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(tokenExpiredBy)).Unix()

	return token.SignedString([]byte(os.Getenv("CUSTOMER_JWT_SECRET_KEY")))
}
