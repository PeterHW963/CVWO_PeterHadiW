import React from "react";
import Logo from "./Logo.tsx";

export default function Header() {
  return (
    <header className="forum-header">
      <div className="logo">
        <a href="/" className="logo-link">
          <Logo />
        </a>
      </div>
      <div className="search-bar">
        {/* Add your search bar component here */}
      </div>
      <div className="user-info">
        <span className="username">{username}</span>
        <img className="/user_avatar.png" alt="image not found" />
      </div>
    </header>
    // <div>
    //   {/* <Logo key={1} /> */}
    //   <Account username="peter" />
    //   <a href="/">test</a>
    // </div>
  );
}
