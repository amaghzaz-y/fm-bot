import { Title } from "solid-start";
import { AiOutlineSend } from "solid-icons/ai";
import Chatbar from "~/components/chatbar";
import Replycard from "~/components/replycard";
export default function Home() {
  return (
    <main>
      <Title>FM-BOT</Title>
      <h1 class="text-center">FM-GPT</h1>
      <Chatbar />
      <Replycard />
    </main>
  );
}
