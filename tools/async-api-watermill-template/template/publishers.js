const {File} = require('@asyncapi/generator-react-sdk');
import {Publishers} from '../components/publishers';

export default async function ({asyncapi, params}) {
    console.log(params.mode)
    if (params.mode !== "client") {
        return;
    }

    let channelWithPublishers = {}

    Object.entries(asyncapi.channels()).forEach(([channelName, channel]) => {
        if (channel.hasPublish()) {
            channelWithPublishers[channelName] = channel
        }
    });
    return (
        <File name="asyncapi_publishers.gen.go">
            <Publishers moduleName={params.moduleName} channels={channelWithPublishers}/>
        </File>
    );
}