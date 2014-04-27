package core

import (
    "github.com/twinj/uuid"
)

type StoryCharacter struct {
    ID uuid.UUID
    Name string
}

type StoryEventTag int

const (
    /* Something is happening within the story that does not fall into the
     * other categories listed.
     */
    STORY_EVENT_NARRATION StoryEventTag = iota
    
    /* A character is saying something.
     * 
     * Example:
     *  Hero says Hi!
     */
    STORY_EVENT_DIALOGUE
    
    /* A character is doing something.
     *
     * Example:
     *  Hero swings his sword!
     */
    STORY_EVENT_ACTION
    
    /* The setting, environment, or background is changing. 
     * 
     * Example:
     *  Meanwhile in Britain...
     */
    STORY_EVENT_TRANSITION
    
    /* Non-digetic discussion between writers. (Aka out-of-character discussion)
     */
    STORY_EVENT_CHAT
)

type StoryEvent struct {
    ID uuid.UUID
    Type StoryEventTag
}

type StorySession struct {
    ID uuid.UUID
    Name string
    Language string
}