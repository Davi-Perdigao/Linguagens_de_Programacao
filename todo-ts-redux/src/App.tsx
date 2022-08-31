import { Heading, IconButton, VStack, useColorMode, Link, Flex } from "@chakra-ui/react";
import { FaSun, FaMoon, FaGithub, FaLinkedin, FaInstagram, FaTwitter, FaFacebook } from 'react-icons/fa';
import TaskList from "./components/Task/List";
import { store } from "./app/store";


function App() {
    const { colorMode, toggleColorMode } = useColorMode();
    store.subscribe(() => {
        localStorage.setItem('tasks', JSON.stringify(store.getState().tasksWatch.tasks));
        localStorage.setItem('tab', store.getState().tabWatch.tab);
    })

    return(
        <VStack p={4} minH='100vh' pb={28}>
            <IconButton 
                aria-label='Trocar tema'
                icon={colorMode === 'light' ? <FaSun /> : <FaMoon />}
                isRound={true} 
                size='md'
                alignSelf='flex-end'
                onClick={toggleColorMode}
            />

            <Heading
                p='5'
                fontWeight='extrabold'
                size='xl'
                bgGradient='linear(to-l, teal.300, blue.500)'
                bgClip='text'
            >
                Lista de tarefas
            </Heading>

            <TaskList />

        </VStack>
    );
}

export default App;