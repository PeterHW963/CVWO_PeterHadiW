import axios from "axios";
import APILINK from "../env.tsx";
import Thread from "../types/Thread";

export default async function AddThread(thread: Thread, jwt: string) {
  const data = {
    stringToken: jwt,
    name: thread.name,
    description: thread.description,
  };
  const response = await axios.post(APILINK + "/thread/create", data);
  if (response.data == "couldnt get cookie") {
    return -1;
  } else if (response.data == "Token Parsing Failed") {
    return -2;
  } else if (response.data == "Token expired") {
    return -3;
  }
  return response.data;
}
