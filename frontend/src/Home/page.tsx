import Navbar from "../components/Navbar/navbar"
import {useState,useEffect} from 'react';
import {CiStar} from "react-icons/ci"

function Home(){
    const [homeMessage,setHomeMessage] = useState<{ [key: string]: any }>({})

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
    },[]);


    return (
        <>
            <Navbar/>
            <div className="mx-5 mt-2">
                <h1 className="text-4xl sm:text-5xl md:text-6xl font-extrabold tracking-tight leading-tight font-sans m-0">
                    <span className="bg-clip-text text-transparent bg-gradient-to-r from-indigo-400 via-purple-400 to-pink-400">Popular</span>
                    <span className="ml-3 text-white/90">Books</span>
                </h1>
            </div>
            <div className =" m-5 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                {
                    Object.keys(homeMessage).map((id) => {
                        const item = homeMessage[id];
                        return (
                          <article key={id} className="bg-white rounded-lg shadow-md hover:shadow-xl hover:scale-[1.02] transition-transform duration-200 p-4 flex flex-col">
                            {item?.CoverImage && (
                              <img src={item.CoverImage} alt={item.Title ?? 'cover'} className="w-full h-40 object-cover rounded-md mb-3" />
                            )}
                            <div className="flex-1">
                              <h3 className="text-xl font-semibold text-gray-900">{item.Title ?? 'Untitled'}</h3>
                              <p className="text-sm text-gray-600 mt-1">By {item.Author ?? 'Unknown'}</p>
                              {item.Description && (
                                <p className="mt-3 text-sm text-gray-700 line-clamp-3">{item.Description}</p>
                              )}
                            </div>
                            <div className="mt-4 flex items-center justify-between">
                              <span className="text-xs text-gray-500">{item.Year ?? ''} {item.Pages ? `• ${item.Pages} pages` : ''}</span>
                              <div className="flex items-center gap-2">
                                <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-yellow-50 text-yellow-700 font-medium" aria-label={`Rating ${item.Rating}`}>
                                  <CiStar className="w-4 h-4 text-yellow-500" />
                                  <span className="text-sm">{Number(item.Rating ?? 0).toFixed(1)}</span>
                                </div>
                              </div>
                            </div>
                          </article>
                        )
                    })
                }                
            </div>
        </>
    )
}

export default Home