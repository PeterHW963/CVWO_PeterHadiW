import React, { useEffect, useState } from "react";
import GetUser from "../../api-calls/Get-User.tsx";
import GetCookie from "../../cookie/Get-Cookie.tsx";
import User from "../../types/User.tsx";
import Account from "./Account/Account.tsx";
import Logo from "./Logo/Logo.tsx";

export default function Header() {
  const [user, setuser] = useState<User>();
  // const [username, setusername] = useState<string>("guest");
  async function getUser() {
    const jwt = GetCookie("jwtauth");
    const data = await GetUser(jwt);
    if (data !== -1 && data !== undefined) {
      console.log(data);
      setuser(data);
      // setusername(data.username);
    }
  }

  useEffect(() => {
    getUser();
  }, []);

  return (
    <header className="forum-header">
      <Logo />

      <div className="search-bar">
        {/* Add your search bar component here */}
      </div>
      {user ? (
        <Account username={user.username} />
      ) : (
        <Account username="guest" />
      )}
      {/* <Account username={username} /> */}
    </header>
    // <div>
    //   {/* <Logo key={1} /> */}
    //   <Account username="peter" />
    //   <a href="/">test</a>
    // </div>
  );
}
