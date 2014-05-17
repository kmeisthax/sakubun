package core

import (
    "github.com/gorilla/sessions"
    "github.com/twinj/uuid"
    "code.google.com/p/go.crypto/bcrypt"
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

/* Represents information needed to authenticate with a password.
 */
type PasswordCredential struct {
    /* The user that is tied to this credential.
     */
    Id uuid.UUID
    
    /* Hash of the password needed to log in.
     */
    PassHash []byte
}

/* Return the ID of the user whose permissions are being used for this request.
 * 
 * The effective UID indicates what set of permissions authorize a particular
 * action within a system. Thus it should be used for permission checks and
 * identification. Note that the effective UID should not be used for 
 */
func GetEffectiveUID (s *sessions.Session) uuid.UUID {
    return s.Values["EUID"].(uuid.UUID)
}

/* Return the ID of the user that logged in to this session.
 * 
 * This authenticated UID indicates who is performing a particular action.
 * It should be used primarily for logging and determining what effective UIDs
 * a user is allowed to assume.
 */
func GetAuthenticatedUID (s *sessions.Session) uuid.UUID {
    return s.Values["AUID"].(uuid.UUID)
}

/* Login the current session as a particular user.
 * 
 * IsMasquerading indicates if a user is assuming another user's credentials.
 * If true, only the Effective user will be set, not the Authenticated user.
 * You should ensure that the Authenticated UID has permission to masquerade as
 * the targeted UID.
 */
func SetLoginSession (s *sessions.Session, u User, IsMasquerading bool) {
    s.Values["EUID"] = u.Id
    
    if (!IsMasquerading) {
        s.Values["AUID"] = u.Id
    }
}

func (u PasswordCredential) CheckPassword(password []byte) bool {
    if (bcrypt.CompareHashAndPassword(u.PassHash, password) == nil) {
        return true
    } else {
        return false
    }
}