import PanelBot from "./panel_bot";
import PanelTop from "./panel_top";

export const PanelLeft = () => {
    return (
        <div id="panel-left" className="h-screen w-96 bg-gray-700 flex flex-col p-4 divide-y divide-gray-500">
            <PanelTop></PanelTop>
            <PanelBot></PanelBot>
        </div>
    )
};

export default PanelLeft;