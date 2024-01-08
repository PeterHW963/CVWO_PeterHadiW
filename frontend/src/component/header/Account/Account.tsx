import React from "react";
import "./Account.css";

interface AccountInfo {
  username: string;
}

export default function Account(props: AccountInfo) {
  const { username } = props;

  return (
    <div>
      <div>{username}</div>
      <div className="account">
        <img src="/user_avatar.png" alt="image not found" />
      </div>
    </div>
  );
}
