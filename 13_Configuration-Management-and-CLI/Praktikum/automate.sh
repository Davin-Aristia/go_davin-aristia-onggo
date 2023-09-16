list_of_my_friends="https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt"
host=$(uname -a)

ROOT="$1 at $(date)"

mkdir "$ROOT"
mkdir "$ROOT/about_me"
mkdir "$ROOT/about_me/personal"
mkdir "$ROOT/about_me/professional"
mkdir "$ROOT/my_friends"
mkdir "$ROOT/my_system_info"

echo "https://www.facebook.com/$2" > "$ROOT/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "$ROOT/about_me/professional/linkedin.txt"
curl "$list_of_my_friends" > "$ROOT/my_friends/list_of_my_friends.txt"
echo "My username: $(whoami)" > "$ROOT/my_system_info/about_this_laptop.txt"
echo "With host: $host" >> "$ROOT/my_system_info/about_this_laptop.txt"
ping -c 3 google.com > "$ROOT/my_system_info/internet_connection.txt"