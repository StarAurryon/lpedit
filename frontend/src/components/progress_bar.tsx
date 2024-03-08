interface Props {
    className?: string;
    textClassName?: string;
    progress: number;
};

export const ProgressBar = ({className, textClassName, progress}: Props) => {
    return (
        <div className={`bg-gray-600 h-10 rounded-md grid-rows-1 grid-cols-1 relative ${className}`}>
            <div className="absolute h-full w-full flex">
                <span className={`block text-center m-auto text-white ${textClassName}`}>{`${progress}%`}</span>
            </div>
            <div className={`rounded-md bg-blue-400 h-full ${className}`} style={{'width': `${progress}%`}}></div>
        </div>
    )
}

export default ProgressBar;