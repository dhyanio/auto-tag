package cloud

import (
	cons "github.com/dhyanio/auto-tag/constant"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func AssumeRole(role, region string) (*session.Session, error) {
	rootSess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(cons.AWSROOTAccess, cons.AWSRootSecret, ""),
	})

	if err != nil {
		return nil, err
	}

	svc := sts.New(rootSess)
	sessionName := cons.ProjectSession
	assumeResult, err := svc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         &role,
		RoleSessionName: &sessionName,
	})

	if err != nil {
		return nil, err
	}

	assumeSess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(*assumeResult.Credentials.AccessKeyId, *assumeResult.Credentials.SecretAccessKey, *assumeResult.Credentials.SessionToken),
	})

	if err != nil {
		return nil, err
	}

	return assumeSess, nil
}
