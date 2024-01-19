import axios from "axios";
import APILINK from "../env.tsx";
// import User from "../types/User";

export default async function GetUser(jwt: string) {
  const response = await axios.post(APILINK + "authenticate", {
    stringToken: jwt,
  });
  if (response.data == "couldnt get cookie") {
    return -1;
  }
  return response.data;
}
