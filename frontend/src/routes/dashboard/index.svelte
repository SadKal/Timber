<script lang="ts">
    import ChatDashboard from "./ChatDashboard.svelte"
    import ChatMain from "./ChatMain.svelte"

    const chats = [
        { user: "PEPE", lastMessage: "Hola" },
        { user: "MIGUEL", lastMessage: "Hola" },
        { user: "ROSA", lastMessage: "ooo i, oo a a, rim ram, ramalamadindon ooo i, oo a a, rim ram, ramalamadindon aaaaaaaaaaaaaaaaaaaaaaaaa" },
        { user: "JUAN", lastMessage: "Hey there!" },
        { user: "ANA", lastMessage: "What's up?" },
        { user: "CARLOS", lastMessage: "How are you doing?" },
        { user: "MARIA", lastMessage: "Long time no see!" },
        { user: "DAVID", lastMessage: "I'm bored." },
        { user: "LAURA", lastMessage: "Did you watch the game last night?" },
        { user: "PEDRO", lastMessage: "Can you help me with something?" },
        { user: "SOFIA", lastMessage: "I love this weather!" },
        { user: "JORGE", lastMessage: "Guess what happened today." },
        { user: "ELENA", lastMessage: "I'm so tired." },
        { user: "PABLO", lastMessage: "Let's grab lunch tomorrow." },
        { user: "LUCIA", lastMessage: "I miss you!" },
        { user: "ALBERTO", lastMessage: "Have you seen the news?" },
        { user: "BEATRIZ", lastMessage: "Can't wait for the weekend!" },
        { user: "SERGIO", lastMessage: "I have a surprise for you." },
        { user: "IRENE", lastMessage: "Remember that trip we took?" },
        { user: "JULIA", lastMessage: "I'm so excited for the concert!" }
    ];
</script>


<div class="w-full h-screen bg-leaf-600 flex justify-between">
    <ChatMain />
    <ChatDashboard {chats} />
</div>