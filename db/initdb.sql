CREATE TABLE pictures
(
    picture_id      serial        not null unique,
    copyright       varchar(255)  not null,
    date            date          not null unique,
    explanation     varchar(5000) not null,
    hd_url          varchar(255)  not null,
    media_type      varchar(255)  not null,
    service_version varchar(255)  not null,
    title           varchar(255)  not null,
    url             varchar(255)  not null,
    PRIMARY KEY (picture_id)
);

CREATE TABLE album
(
    id           serial not null unique,
    requested_at date   not null,
    picture_id   int    not null,
    PRIMARY KEY (id),
    FOREIGN KEY (picture_id) REFERENCES pictures (picture_id)
);

INSERT INTO pictures
VALUES (DEFAULT, 'ALSJ', '2022-07-30',
        'Get out your red/blue glasses and check out this stereo view from lunar orbit. The 3D anaglyph was created from two photographs (AS11-44-6633, AS11-44-6634) taken by astronaut Michael Collins during the 1969 Apollo 11 mission. It features the lunar module ascent stage, dubbed The Eagle, rising to meet the command module in lunar orbit on July 21. Aboard the ascent stage are Neil Armstrong and Buzz Aldrin, the first to walk on the Moon. The smooth, dark area on the lunar surface is Mare Smythii located just below the equator on the extreme eastern edge of the Moon''s near side.  Poised beyond the lunar horizon is our fair planet Earth.',
        'https://apod.nasa.gov/apod/image/2207/AS11JK44-6633-34.jpg', 'image', 'v1', 'The Eagle Rises',
        'https://apod.nasa.gov/apod/image/2207/AS11JK44-6633-34_1100px.jpg');

INSERT INTO pictures
VALUES (DEFAULT, 'Ian Griffin', '2022-07-29',
        'SOFIA, the Stratospheric Observatory for Infrared Astronomy, is a Boeing 747SP aircraft modified to carry a large reflecting telescope into the stratosphere. The ability of the airborne facility to climb above about 99 percent of Earth''s infrared-blocking atmosphere has allowed researchers to observe from almost anywhere over the planet. On a science mission flying deep into the southern auroral oval, astronomer Ian Griffin, director of New Zealand\u2019s Otago Museum, captured this view from the observatory''s south facing starboard side on July 17. Bright star Canopus shines in the southern night above curtains of aurora australis, or southern lights. The plane was flying far south of New Zealand at the time at roughly 62 degrees southern latitude. Unfortunately, after a landing at Christchurch severe weather damaged SOFIA requiring repairs and the cancellation of the remainder of its final southern hemisphere deployment.',
        'https://apod.nasa.gov/apod/image/2207/ASC05954-Edit.jpg', 'image', 'v1', 'SOFIA''s Southern Lights',
        'https://apod.nasa.gov/apod/image/2207/ASC05954-Edit1024.jpg');

INSERT INTO album
VALUES (DEFAULT, '2022-07-30', 1);

INSERT INTO album
VALUES (DEFAULT, '2022-07-30', 2);
