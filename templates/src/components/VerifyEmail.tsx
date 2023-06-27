function VerifyEmail() {
    const statusCode = 404

    return (
        <div>
            {
                statusCode === 404 ?
                    <>
                        <div className="notification-error error">
                            <h2>Verification Unsuccessful</h2>
                            <p>We were unable to verify your account. Please try again later.</p>
                        </div>
                    </>
                    :
                    <>
                        <div className="notification-success">
                            <h2>Verification Successful</h2>
                            <p>Your account has been successfully verified.</p>
                        </div>
                    </>
            }
        </div>
    );
}

export default VerifyEmail;