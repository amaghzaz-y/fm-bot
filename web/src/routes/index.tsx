import { Title } from "solid-start";
import { AiOutlineSend } from "solid-icons/ai";
import Chatbar from "~/components/chatbar";
import Replycard from "~/components/replycard";
export default function Home() {
  return (
    <main>
      <Title>FM-BOT</Title>
      <h1 class="text-center">FM-GPT</h1>
      <div class="flex flex-col gap-2 p-2">
        <Chatbar />
        <Replycard />
      </div>
    </main>
  );
}
