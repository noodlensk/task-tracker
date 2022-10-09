const {File} = require('@asyncapi/generator-react-sdk');
import {Subscribers} from '../components/subscribers';

export default async function ({asyncapi, params}) {
    if (params.mode !== "server") {
        return;
    }

    let channelWithSubscribers = {}

    Object.entries(asyncapi.channels()).forEach(([channelName, channel]) => {
        if (channel.hasSubscribe()) {
            channelWithSubscribers[channelName] = channel
        }
    });
    return (
        <File name="asyncapi_subscribers.gen.go">
            <Subscribers moduleName={params.moduleName} channels={channelWithSubscribers}/>
        </File>
    );
}
module.exports = {
    'generate:after': generator => console.log('This runs after generation is complete')
}