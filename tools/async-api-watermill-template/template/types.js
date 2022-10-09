const {File} = require('@asyncapi/generator-react-sdk');
import {Types} from '../components/types';

export default async function ({asyncapi, params}) {
    let messages = []

    Object.entries(asyncapi.channels()).forEach(([channelName, channel]) => {
        if (params.mode === "server" && channel.hasSubscribe()) {
            messages.push(channel.subscribe().message())
        } else if (params.mode === "client" && channel.hasPublish()) {
            messages.push(channel.publish().message())
        }
    });

    return (
        <File name="asyncapi_types.gen.go">
            <Types moduleName={params.moduleName} messages={messages}/>
        </File>
    );
}