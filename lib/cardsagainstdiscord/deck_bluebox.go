package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "bluebox",
		Description: "Blue box expansion",
		Prompts: []*PromptCard{
			{Prompt: `%s may pass, but %s will last forever.`},
			{Prompt: `%s will never be the same after %s.`},
			{Prompt: `"This is madness." "No, THIS IS %s!"`},
			{Prompt: `2 AM in the city that never sleeps. The door swings open and she walks in, legs up to here. Something in her eyes tells me she's looking for %s.`},
			{Prompt: `Adventure. Romance. %s. From Paramount Pictures, "%s."`},
			{Prompt: `And today's soup is Cream of %s.`},
			{Prompt: `And would you like those buffalo wings mild, hot, or %s?`},
			{Prompt: `Armani suit: $1,000. Dinner for two at that swanky restaurant: $300. The look on her face when you surprise her with %s: priceless.`},
			{Prompt: `As king, how will I keep the peasants in line? %s`},
			{Prompt: `Behind every powerful man is %s.`},
			{Prompt: `Come to Dubai, where you can relax in our world-famous spas, experience the nightlife, or simply enjoy %s by the poolside.`},
			{Prompt: `Dammit Gary. You can't just solve every problem with %s.`},
			{Prompt: `Dear Leader Kim Jong-un, our village praises your infinite wisdom with a humble offering of %s.`},
			{Prompt: `Do not fuck with me! I am literally %s right now.`},
			{Prompt: `Do the Dew with our most extreme flavor yet! Get ready for Mountain Dew %s!`},
			{Prompt: `Do you lack energy? Does it sometimes feel like the whole world is %s ? Ask your doctor about Zoloft.®`},
			{Prompt: `Don't forget! Beginning this week, Casual Friday will officially become "%s Friday."`},
			{Prompt: `Don't worry kid. It gets better. I've been living with %s for 20 years.`},
			{Prompt: `Every step towards %s gets me a little bit closer to %s.`},
			{Prompt: `Everybody join hands and close your eyes. Do you sense that? That's the presence of %s in this room.`},
			{Prompt: `Forget everything you know about %s, because now we've supercharged it with %s!`},
			{Prompt: `Get ready for the movie of the summer! One cop plays by the book. The other's only interested in one thing: %s.`},
			{Prompt: `Having the worst day EVER. #%s`},
			{Prompt: `Heed my voice, mortals! I am the god of %s , and I will not tolerate %s!`},
			{Prompt: `Help me doctor, I've got %s in my butt!`},
			{Prompt: `Here at the Academy for Gifted Children, we allow students to explore %s at their own pace.`},
			{Prompt: `Hi MTV! My name is Kendra, I live in Malibu, I'm into %s, and I love to have a good time.`},
			{Prompt: `Hi, this is Jim from accounting. We noticed a $1,200 charge labeled "%s." Can you explain?`},
			{Prompt: `Honey, I have a new role-play I want to try tonight! You can be %s, and I'll be %s.`},
			{Prompt: `How am I compensating for my tiny penis?%s`},
			{Prompt: `I am become %s, destroyer of %s!`},
			{Prompt: `I don't mean to brag, but they call me the Micheal Jordan of %s.`},
			{Prompt: `I have a strict policy. First date, dinner. Second date, kiss. Third date, %s.`},
			{Prompt: `I work my ass off all day for this family, and this is what I come home to? %s!?`},
			{Prompt: `I'm Miss Tennessee, and if I could make the world better by changing one thing, I would get rid of %s.`},
			{Prompt: `I'm pretty sure I'm high right now, because I"m absolutely mesmerized by %s.`},
			{Prompt: `I'm sorry, Mrs. Chen, but there was nothing we could do. At 4:15 this morning, your son succumbed to %s.`},
			{Prompt: `I'm sorry, sir, but we don't allow %s at the country club.`},
			{Prompt: `If you can't handle %s, you'd better stay away from %s.`},
			{Prompt: `If you had to describe the Card Czar, using only one of the cards in your hand, which one would it be? %s`},
			{Prompt: `In his farewell address, George Washington famously warned Americans about the dangers of %s.`},
			{Prompt: `In his new action comedy, Jackie Chan must fend off ninjas while also dealing with %s.`},
			{Prompt: `In return for my soul, the Devil promised me %s, but all I got was %s.`},
			{Prompt: `In the beginning, there was %s. And the Lord said, "Let there be %s."`},
			{Prompt: `It lurks in the night. It hungers for flesh. This summer, no one is safe from %s.`},
			{Prompt: `James is a lonely boy. But when he discovers a secret door in his attic, he meets a magical new friend: %s.`},
			{Prompt: `Life's pretty tough in the fast lane. That's why I never leave the house without %s.`},
			{Prompt: `Listen Gary, I like you. But if you want that corner office, you're going to have to show me %s.`},
			{Prompt: `Man, this is bullshit. Fuck %s.`},
			{Prompt: `My grandfather worked his way up from nothing. When he came to this country, all he had was the shoes on his feet and %s.`},
			{Prompt: `Now in bookstores: "The Audacity of %s" by Barack Obama.`},
			{Prompt: `Oprah's book of the month is "%s For %s: A Story of Hope."`},
			{Prompt: `Patient presents with %s. Likely a result of %s.`},
			{Prompt: `Puberty is a time of change. You might notice hair growing in new places. You might develop an intrest in %s. This is normal.`},
			{Prompt: `She's up all night for good fun. I'm up all night for %s.`},
			{Prompt: `The Japanese have developed a smaller, more efficient version of %s.`},
			{Prompt: `The six things I could never do without: oxygen, Facebook, chocolate, Netflix, friends, and %s LOL!`},
			{Prompt: `This is America. If you don't work hard, you don't succeed. I don't care if you're black, white, purple, or %s.`},
			{Prompt: `This is the prime of my life. I'm young, hot, and full of %s.`},
			{Prompt: `This year's hottest album is "%s" by %s.`},
			{Prompt: `To become a true Yanomami warrior, you must prove that you can withstand %s without crying out.`},
			{Prompt: `Tonight we will have sex. And afterwards, If you'd like, a little bit of %s.`},
			{Prompt: `We never did find %s, but along the way we sure learned a lot about %s.`},
			{Prompt: `Well if %s is good enough for %s, it's good enough for me.`},
			{Prompt: `Well what do you have to say for yourself, Casey? This is the third time you've been sent to the principal's office for %s.`},
			{Prompt: `Wes Anderson's new film tells the story of a precocious child coming to terms with %s.`},
			{Prompt: `What killed my boner? %s`},
			{Prompt: `What's fun until it gets weird? %s`},
			{Prompt: `What's making things awkward in the sauna? %s and %s`},
			{Prompt: `When I was a kid, we used to play Cowboys and %s.`},
			{Prompt: `WHOOO! God damn I love %s!`},
			{Prompt: `Why am I broke? %s`},
			{Prompt: `Why won't you make love to me anymore? Is it %s?`},
			{Prompt: `Y'all ready to get this thing started? I'm Nick Cannon, and this is America's Got %s.`},
			{Prompt: `Yo' mama's so fat she %s!`},
			{Prompt: `You are not alone. Millions of Americans struggle with %s every day.`},
			{Prompt: `You know, once you get past %s, %s ain't so bad.`},
			{Prompt: `You Won't Believe These 15 Hilarious %s Bloopers!`},
			{Prompt: `You've seen the bearded lady! You've seen the ring of fire! Now, ladies and gentlemen, feast your eyes upon %s!`},
		},
		Responses: []ResponseCard{
			`10 Incredible Facts About the Anus`,
			`40 acres and a mule`,
			`A bass drop so huge it tears the starry vault asunder to reveal the face of God`,
			`A bunch of idiots playing a card game instead of interacting like normal humans`,
			`A buttload of candy`,
			`A chimpanzee in sunglasses fucking your wife`,
			`A constant need for validation`,
			`A crazy little thing called love`,
			`A dance move that's just sex`,
			`A disappointing salad`,
			`A face full of horse cum`,
			`A fart`,
			`A for-real lizard that spits blood from its eyes`,
			`A gender identity that can only be conveyed through slam poetry`,
			`A giant powdery manbaby`,
			`A hopeless amount of spiders`,
			`A horse with no legs`,
			`A kiss on the lips`,
			`A lil’ stupid ass bitch`,
			`A man who is so cool that he rides on a motorcycle`,
			`A manhole`,
			`A mouthful of potato salad`,
			`A Native American who solves crimes by going into the spirit world`,
			`A one-way ticket to Gary, Indiana`,
			`A peyote-fueled vision quest`,
			`A pizza guy who fucked up`,
			`A possible Muslim`,
			`A reason not to commit suicide`,
			`A self-microwaving burrito`,
			`A sex comet from Neptune that plunges the Earth into eternal sexiness`,
			`A sex goblin with a carnival penis`,
			`A shiny rock that proves I love you`,
			`A team of lawyers`,
			`A thrilling chase over the rooftops of Rio de Janeiro`,
			`A turd`,
			`A Ugandan warlord`,
			`A whole lotta woman`,
			`A whole new kind of porn`,
			`A woman`,
			`A zero-risk way to make $2,000 from home`,
			`Actual mutants with medical conditions and no superpowers`,
			`Africa`,
			`AIDS monkeys`,
			`All the single ladies`,
			`All these decorative pillows`,
			`Almost giving money to a homeless person`,
			`Ambiguous sarcasm`,
			`An inability to form meaningful relationships`,
			`An interracial handshake`,
			`An oppressed people with a vibrant culture`,
			`An overwhelming variety of cheeses`,
			`An unforgettable quinceañera`,
			`An uninterrupted history of imperialism and exploitation`,
			`Anal fissures like you wouldn't believe`,
			`Ancient Athenian boy-fucking`,
			`Angelheaded hipsters burning for the ancient heavenly connection to the starry dynamo in the machinery of night`,
			`Ass to mouth`,
			`Backwards knees`,
			`Bathing in moonsblood and dancing around the ancient oak`,
			`Being a terrible mother`,
			`Being John Malkovich`,
			`Being nine years old`,
			`Being paralyzed from the neck down`,
			`Being popular and good at sports`,
			`Being worshipped as the one true God`,
			`Beloved television star Bill Cosby`,
			`Blackface`,
			`Blackula`,
			`Blowjobs for everyone`,
			`Boring Vaginal sex`,
			`Bouncing up and down`,
			`Breastfeeding a ten year old`,
			`Breastfeeding a ten-year-old`,
			`Bullets`,
			`Butt stuff`,
			`Calculating every mannerism so as not to suggest homosexuality`,
			`Cancer`,
			`Changing a person's mind with logic and facts`,
			`Child Protective Services`,
			`Child support payments`,
			`Common-sense gun control legislation`,
			`Cool, relatable cancer teens`,
			`Crazy opium eyes`,
			`Crippling social anxiety`,
			`Crying and shitting and eating spaghetti`,
			`Cute boys`,
			`Cutting off a flamingo's legs with garden shears`,
			`Daddy`,
			`Daddy's credit card`,
			`Deez nuts`,
			`Dem titties`,
			`Depression`,
			`Doing the right stuff to her nipples`,
			`Doo-doo`,
			`Drinking responsibly`,
			`Eating together like a god damn family for once`,
			`Eggs`,
			`Ejaculating inside another man's wife`,
			`Ejaculating live bees and the bees are angry`,
			`Ennui`,
			`Every ounce of charisma left in Mick Jagger's tired body`,
			`Exploding pigeons`,
			`Falling into the toilet`,
			`Figuring out how to have sex with a dolphin`,
			`Filling a man's anus with concrete`,
			`Finally finishing off the Indians`,
			`Free ice cream, yo`,
			`Fresh dill from the patio`,
			`Fucking a corpse back to life`,
			`Generally having no idea what's going on`,
			`Genghis Khan's DNA`,
			`Getting all offended`,
			`Getting caught by the police and going to jail`,
			`Getting down to business to defeat the Huns`,
			`Getting drive-by shot`,
			`Getting eaten alive by Guy Fieri`,
			`Getting offended`,
			`Getting shot by the police`,
			`Getting shot out of a cannon`,
			`Giant sperm from outer space`,
			`Going down on a woman, discovering that her vagina is filled with eyeballs, and being totally into that`,
			`Going inside at some point because of the mosquitoes`,
			`Going to a high school reunion on ketamine`,
			`Grammar nazis who are also regular Nazis`,
			`Growing up chained to a radiator in perpetual darkness`,
			`Gwyneth Paltrow's opinions`,
			`Having been dead for a while`,
			`How awesome I am`,
			`Immortality cream`,
			`Important news about Taylor Swift`,
			`Injecting speed into one arm and horse tranquilizer into the other`,
			`Interspecies marriage`,
			`Irrefutable evidence that God is real`,
			`Jizz`,
			`Kale`,
			`Khakis`,
			`Letting out 20 years’ worth of farts`,
			`Like a million alligators`,
			`Lots and lots of abortions`,
			`Meaningless sex`,
			`Mediocrity`,
			`Moderate-to-severe joint pain`,
			`Mom's new boyfriend`,
			`Morpheus`,
			`My boyfriend's stupid penis`,
			`My dad's dumb fucking face`,
			`My dead son's baseball glove`,
			`My first period`,
			`My sex dungeon`,
			`My worthless son`,
			`Neil Diamond's Greatest Hits`,
			`Never having sex again`,
			`No clothes on, penis in vagina`,
			`No longer finding any Cards Against Humanity card funny`,
			`Not believing in giraffes`,
			`Oil!`,
			`One unforgettable night of passion`,
			`Our new Buffalo Chicken Dippers®!`,
			`Out-of-this-world bazongas`,
			`Owls, the perfect predator`,
			`P.F. Chang himself`,
			`Party Mexicans`,
			`Peeing into a girl's butt to make a baby`,
			`Potato`,
			`Prince Ali, fabulous he, Ali Ababwa`,
			`Pussy`,
			`Rabies`,
			`Reading the entire End-User License Agreement`,
			`Ripping a dog in half`,
			`Robots who just want to party`,
			`Russian super-tuberculosis`,
			`Seeing my village burned and my family slaughtered before my eyes`,
			`Seeing things from Hitler's perspective`,
			`September 11th, 2001`,
			`Setting my balls on fire and cartwheeling to Ohio`,
			`Shapes and colors`,
			`Sharks with legs`,
			`Shitting all over the floor like a bad, bad girl`,
			`Slowly easing down onto a cucumber`,
			`Smoking crack, for instance`,
			`Snorting coke off a clown's boner`,
			`Social justice warriors with flamethrowers of compassion`,
			`Some shit-hot guitar licks`,
			`Some sort of Asian`,
			`Sports`,
			`Storing a bunch of acorns in my pussy`,
			`Stuffing a child's face with Fun Dip® until he starts having fun`,
			`Stupid`,
			`Such a big boy`,
			`Sucking all the milk out of a yak`,
			`Sudden penis loss`,
			`Teaching a girl how to handjob the penis`,
			`Texas`,
			`The Abercrombie & Fitch lifestyle`,
			`The all-new Nissan Pathfinder with 0.9% APR financing!`,
			`The amount of gay I am`,
			`The basic suffering that pervades all of existence`,
			`The best taquito in the galaxy`,
			`The black half of Barack Obama`,
			`The color "puce"`,
			`The complex geopolitical quagmire that is the Middle East`,
			`The dentist`,
			`The eight gay warlocks who dictate the rules of fashion`,
			`The eighth graders`,
			`The euphoric rush of strangling a drifter`,
			`The ghost of Marlon Brando`,
			`The haunting stare of an Iraqi child`,
			`The inability to form meaningful relationships`,
			`The male gaze`,
			`The passage of time`,
			`The peaceful and nonthreatening rise of China`,
			`The power of the Dark Side`,
			`The right amount of cocaine`,
			`The safe word`,
			`The secret formula for ultimate female satisfaction`,
			`The size of my penis`,
			`The sweet song of sword against sword and the braying of mighty war beasts`,
			`The swim team, all at once`,
			`The tiger that killed my father`,
			`The tiniest shred of evidence that God is real`,
			`The unbelievable world of mushrooms`,
			`The white half of Barack Obama`,
			`Three consecutive seconds of happiness`,
			`Throwing stones at a man until he dies`,
			`Too much cocaine`,
			`Total fucking chaos`,
			`Treasures beyond your wildest dreams`,
			`Turning the rivers red with the blood of infidels`,
			`Two whales fucking the shit out of each other`,
			`Unquestioning obedience`,
			`Unrelenting genital punishment`,
			`Unsheathing my massive horse cock`,
			`Vegetarian options`,
			`Walking into a glass door`,
			`Wearing glasses and sounding smart`,
			`Western standards of beauty`,
			`What Jesus would do`,
			`Whatever a McRib® is made of`,
			`Whatever you wish, mother`,
			`Whispering all sexy`,
			`%blank`,
			`%blank`,
		},
	}

	AddPack(pack)
}
