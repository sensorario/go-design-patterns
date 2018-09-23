# Behavioral » Observer

## Description

This pattern uncouple an event from its possible handlers. It is useful to
trigger many actions on same events. And it is useful whenever the number of
action of this event grows.

Key elements of this patterns are the publisher and the subscriber. The
publisher will emit an event. All subscriber that handle that event are called
when the particular event is triggered.

## Implementation


