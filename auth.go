client, err := app.Auth(ctx)
if err != nil {
        log.Fatalf("error getting Auth client: %v\n", err)
}

token, err := client.VerifyIDToken(ctx, idToken)
if err != nil {
        log.Fatalf("error verifying ID token: %v\n", err)
}

log.Printf("Verified ID token: %v\n", token)