import { Title } from "solid-start";
import { AiOutlineSend } from "solid-icons/ai";
import Chatbar from "~/components/chatbar";
import Replycard from "~/components/replycard";
import { Show } from "solid-js";
import { Answer } from "~/store";
export default function Home() {
  return (
    <main>
      <Title>FM-BOT</Title>
      <h1 class="text-center">FM-GPT</h1>
      <div class="flex flex-col gap-2 p-2 justify-center items-center">
        <Chatbar />
        <Show when={Answer().length > 2}>
          <Replycard />
        </Show>
      </div>
    </main>
  );
}
