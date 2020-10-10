<html>
日本語は英文の後にあります。

<p>Hello {{.UserName}},<br>
Today’s lunch group has been decided!</p>

<p>■ Lunch date<br>
{{.Party.Date}}<br>
{{.Party.TimeFrom}}-{{.Party.TimeTo}}</p>

<p>■ Members<br>
{{range $index, $name := .Party.MemberUserNames}}○ {{$name}}<br>
{{end}}
</p>
<!--TODO: Set belongings for each member. 部署名、職種 -->

<!-- TODO: Set meeting point after Admin tool is created.
■ Meeting Point
Rakuten Crimson House 9F Cafeteria Entrance
-->

<p>■ Chat page<br>
You can contact the members on the following page.<br>
  <a href="https://www.mixlunch.com/plan/{{.Party.ChatRoomId}}/">https://www.mixlunch.com/plan/{{.Party.ChatRoomId}}/</a>
</p>

<p>You can access the URL with your cell phone.</p>

<!-- TODO: Generate QR Code URLs and set them in any storage service
[QR Code]
-->

<p>Thanks,<br>
Mixlunch</p>

<br>

<p>{{.UserName}} さん<br>
本日のランチグループが確定しました！</p>

<p>■ ランチの日時<br>
{{.Party.Date}}<br>
{{.Party.TimeFrom}}-{{.Party.TimeTo}}</p>

<p>■ メンバー<br>
{{range $index, $name := .Party.MemberUserNames}}○ {{$name}}<br>
{{end}}
</p>
<!--TODO: Set belongings for each member. 部署名、職種 -->

<!-- TODO: Set meeting point after Admin tool is created.
■ 待ち合わせ場所
Rakuten Crimson House 9F Cafeteria 入り口
-->

<p>■ チャット画面<br>
以下の画面からチャットでやり取りができます。<br>
  <a href="https://www.mixlunch.com/plan/{{.Party.ChatRoomId}}/">https://www.mixlunch.com/plan/{{.Party.ChatRoomId}}/</a>
</p>

<p>スマートフォンからも利用できます。</p>

<!-- TODO: Generate QR Code URLs and set them in any storage service
[QR Code]
-->

Mixlunch
</html>
