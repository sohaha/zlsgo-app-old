package manageBusiness

var (
	avatarPrefix  = "/resource"
	avatarPath    = "." + avatarPrefix + "/static/avatar"
	avatarTemPath = "." + avatarPrefix + "/static/avatar/tmp"

	avatarType       = [2]string{".jpg", ".png"}
	avatarSize int64 = 2048 // k
)
