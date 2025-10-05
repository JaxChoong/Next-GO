import Navbar from "../components/Navbar/navbar"
import {useState,useEffect} from 'react';
function Home(){
    const [homeMessage,setHomeMessage] = useState(0)

    useEffect(() => {
        const fetchData = async () => {
            try{
                const response = await fetch("http://localhost:8080/");
                if (!response.ok){
                    throw new Error("Fetch Home Message error!")
                }
                const message = await response.json()  
                setHomeMessage(message.message)
            }
            catch(error){
                console.log(error)
            }
        };
        fetchData();
    });

    return (
        <>
            <Navbar/>
            {homeMessage}
        </>
    )
}

export default Home