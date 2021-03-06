package core

import (
    "github.com/gorilla/sessions"
    "github.com/twinj/uuid"
    "code.google.com/p/go.crypto/bcrypt"
    
    "time"
)

func getuuid(uuidstring string) uuid.UUID {
    id, err := uuid.ParseUUID(uuidstring)
    if (err != nil) {
        panic(err)
    }
    
    return id
}

/* UUID representing that the user is logged out. */
var Anonymous uuid.UUID = getuuid("{00000000-0000-4000-8000-00000000A404}")

/* A user identity managed by Sakubun.
 */
type User struct {
    Id uuid.UUID
    
    /* The name of a user within a system. This should be in the format of a
     * valid e-mail address, but it may or may not be the user's contact e-mail
     */
    Username string
}

/* Information relating to authenticated access. (e.g. a login session)
 * 
 * Authenticated access is granted by hitting the authentication endpoint with
 * appropriate data to authorize access to a particular UID. Once any relevant
 * credentials for an authentication have been validated, the user recieves
 * an authentication token to be delivered in future authenticated requests
 * (e.g. by session cookie, HTTP Bearer authentication, etc)
 */
type Authentication struct {
    /* ID of the authentication event. */
    Id uuid.UUID
    
    /* Current nonce required to use access granted by this authentication. */
    Token []byte
    
    /* The UUID that was authenticated. */
    AuthenticatedUID uuid.UUID
    
    /* The credential used to authenticate the user.
     * 
     * This credential may contain additional access limitations, which should
     * be taken into consideration when checking permissions.
     */
    CredentialID uuid.UUID
}

/* Represents an abstract set of credentials to be used to authenticate a user.
 * The actual credential data is stored in separate structures.
 */
type Credential struct {
    /* The date of the credential. */
    Id uuid.UUID
    
    /* Time that the credential was issued.
     * 
     * Credentials may not be used before their issue date.
     */
    IssuedOn time.Time
    
    /* Time that the credential is no longer valid after.
     * 
     * Credentials after this date are considered Expired and should not be
     * used.
     * 
     * If a credential's expiration date is zero (see IsZero method) then the
     * credential is said to not expire.
     */
    ExpiresOn time.Time
    
    /* Indicates if a credential is revoked and should not be used anymore.
     * 
     * Credentials that are used after revocation should be logged and
     * considered a security incident.
     */
    IsRevoked bool
    
    /* Indicates the type of credential data.
     * 
     * The validation process for a particular credential is determined by the
     * type of credential 
     */
    CredentialType string
}

/* Represents a credential type that allows users to authenticate with a text
 * string, called a password.
 * 
 * For security reasons the actual password is enciphered with a secure
 * password hashing algorithm designed specifically to frustrate brute-force
 * attacks in the event that the credentials database is compromised.
 * 
 * (Note that the existence of cryptocurrencies based on a proof-of-work system
 * utilizing secure password hashing algorithms means that there is now
 * economic incentive to develop dedicated hardware to brute-force password
 * hashes. Therefore the security of passwords stored in this method is bounded
 * by the price of Litecoin mining ASICs.)
 */
type PasswordCredential struct {
    /* The user that is tied to this credential.
     */
    Id uuid.UUID
    
    /* Hash of the password needed to authenticate.
     */
    PassHash []byte
}

/* Return the ID of the user that logged in to this session.
 */
func GetAuthenticatedUID (s *sessions.Session) uuid.UUID {
    return s.Values["UserID"].(uuid.UUID)
}

/* Called after a credential is validated.
 * 
 * An authentication is either updated or created for the particular user-
 * credential pair, depending on if the session already posses auth data.
 * 
 * In either case, the session is guaranteed to contain a valid authentication
 * ID and token.
 */
func ApplySessionAuthentication (s *sessions.Session, u User, c Credential) {
    s.Values["UserID"] = u.Id
    
    if (s.Values["AuthID"] != nil) {
        //FIXME: Refresh the user's authentication tokens.
    } else {
        //FIXME: Create new authentication tokens.
    }
}

func (u PasswordCredential) CheckPassword(password []byte) bool {
    if (bcrypt.CompareHashAndPassword(u.PassHash, password) == nil) {
        return true
    } else {
        return false
    }
}