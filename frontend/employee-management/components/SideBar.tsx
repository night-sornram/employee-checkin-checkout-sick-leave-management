'use client'
 
import { usePathname  , useRouter } from 'next/navigation'
import { EnterIcon , ExitIcon , CounterClockwiseClockIcon , CalendarIcon , CheckIcon , PersonIcon  } from "@radix-ui/react-icons"
import { Button } from "@/components/ui/button"
import { useSession } from 'next-auth/react'


export default function SideBar({children}: {children: React.ReactNode}) {
    const params = usePathname()
    const router = useRouter()
    const { data: session } = useSession()

    return (
        <div className=" h-[93vh] max-w-screen flex flex-row">
            <div className=" w-[15vw] h-full border-r flex flex- p-5">
                <div className=" flex flex-col space-y-5 w-full">
                    <h1 className="text-sm">
                        CHECK-IN CHECK-OUT
                    </h1>
                    <div className=" flex flex-col space-y-3 w-full">
                        <Button className={` ${ params === "/attendance/checkin" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : ""}   `} onClick={()=>{router.push("/attendance/checkin")}} variant="ghost">
                            <EnterIcon className="mr-2 h-5 w-5"/>
                            <h1 >
                                Check In
                            </h1> 
                        </Button>
                        <Button className={` ${ params === "/attendance/checkout" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : ""}   `} onClick={()=>{router.push("/attendance/checkout")}} variant="ghost">
                            <ExitIcon  className="mr-2 h-5 w-5" />
                            <h1>
                                Check Out
                            </h1> 
                        </Button>
                        <Button className={` ${ params === "/attendance/history" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : ""}   `} onClick={()=>{router.push("/attendance/history")}} variant="ghost">
                            <CounterClockwiseClockIcon  className="mr-2 h-5 w-5" />
                            <h1 >
                                History
                            </h1> 
                        </Button>
                    </div>
                    <h1 className="text-sm">
                        LEAVE MANAGEMENT
                    </h1>
                    <div className=" flex flex-col space-y-3 w-full">
                        <Button className={` ${ params === "/leave-request/request" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : ""}   `} onClick={()=>{router.push("/leave-request/request")}} variant="ghost">
                            <CalendarIcon className="mr-2 h-5 w-5" />
                            <h1 >
                                Request
                            </h1> 
                        </Button>
                        <Button className={` ${ params === "/leave-request/history" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : " hover:"}   `} onClick={()=>{router.push("/leave-request/history")}} variant="ghost">
                            <CounterClockwiseClockIcon  className="mr-2 h-5 w-5" />
                            <h1 >
                                History
                            </h1> 
                        </Button>
                    </div>
                    {
                        session?.user.role === "admin" ? (
                            <div className=" flex flex-col space-y-5 w-full">
                                <h1 className="text-sm">
                                    ADMIN DASHBOARD
                                </h1>
                                <div className=" flex flex-col space-y-3 w-full">
                                    <Button className={` ${ params === "/dashboard/create-employee" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : ""}   `} onClick={()=>{router.push("/dashboard/create-employee")}} variant="ghost">
                                        <PersonIcon className="mr-2 h-5 w-5" />
                                        <h1 >
                                            Create Employee
                                        </h1> 
                                    </Button>
                                    <Button className={` ${ params === "/dashboard/approve-leave" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : " hover:"}   `} onClick={()=>{router.push("/dashboard/approve-leave")}} variant="ghost">
                                        <CheckIcon  className="mr-2 h-5 w-5" />
                                        <h1 >
                                            Approve Leave
                                        </h1> 
                                    </Button>
                                </div>
                            </div>
                        )
                        :
                        null
                    }
                    
                    
                </div>
           </div>
            {children}
        </div>
    
    
    )}