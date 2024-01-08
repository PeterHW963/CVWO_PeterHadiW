import React from "react";
import Account from "./Account/Account.tsx";

export default function Header() {
  return (
    <div>
      {/* <Logo key={1} /> */}
      <Account username="peter" />
      <a href="/">test</a>
    </div>
  );
}
