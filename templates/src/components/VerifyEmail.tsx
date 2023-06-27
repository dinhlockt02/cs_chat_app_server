function VerifyEmail() {
    const statusCode = 404

    return (
        <div>
            {
                statusCode === 404 ?
                    <div className="msg msg-error z-depth-3">Verify email failed. Please try again </div>
                    :
                    <div className="msg msg-info z-depth-3">You're verify your email successful </div>
            }
        </div>
    );
}

export default VerifyEmail;