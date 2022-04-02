CREATE TABLE public.news (
  guid character varying(255) PRIMARY KEY,
  title character varying(255) NOT NULL,
  content text,
  created_at timestamp NOT NULL,
  deleted_at timestamp,
  updated_at timestamp,
  published_at timestamp
);

CREATE TABLE public.tags (
  guid character varying(255) PRIMARY KEY,
  name character varying(255) UNIQUE,
  created_at timestamp NOT NULL,
  deleted_at timestamp,
  updated_at timestamp
);

CREATE TABLE public.news_tags (
  guid character varying(255) PRIMARY KEY,
  guid_news character varying(255),
  guid_tag character varying(255),
  created_at timestamp NOT NULL,
  deleted_at timestamp,
  updated_at timestamp
);

CREATE TABLE public.topics (
  guid character varying(255) PRIMARY KEY,
  name character varying(255) UNIQUE,
  created_at timestamp NOT NULL,
  deleted_at timestamp,
  updated_at timestamp
);

CREATE TABLE public.topic_news (
  guid character varying(255) PRIMARY KEY,
  guid_topic character varying(255),
  guid_news character varying(255),
  created_at timestamp NOT NULL,
  deleted_at timestamp,
  updated_at timestamp
);

ALTER TABLE public.news_tags
    ADD CONSTRAINT guid_news_fkey FOREIGN KEY (guid_news) REFERENCES public.news(guid);

ALTER TABLE public.news_tags
    ADD CONSTRAINT guid_tag_fkey FOREIGN KEY (guid_tag) REFERENCES public.tags(guid);

ALTER TABLE public.topic_news
    ADD CONSTRAINT guid_topic_fkey FOREIGN KEY (guid_topic) REFERENCES public.topics(guid);

ALTER TABLE public.topic_news
    ADD CONSTRAINT guid_news_fkey FOREIGN KEY (guid_news) REFERENCES public.news(guid);

INSERT INTO public.news(guid, title, content, created_at) VALUES('n1','News 1','Lorem ipsum dolor sit amet',NOW());
INSERT INTO public.news(guid, title, content, created_at) VALUES('n2','News 2','Lorem ipsum dolor sit amet',NOW());
INSERT INTO public.news(guid, title, content, created_at) VALUES('n3','News 3','Lorem ipsum dolor sit amet',NOW());

INSERT INTO public.tags(guid, name, created_at) VALUES('t1','Tags 1',NOW());
INSERT INTO public.tags(guid, name, created_at) VALUES('t2','Tags 2',NOW());

INSERT INTO public.news_tags(guid, guid_news, guid_tag, created_at) VALUES('nt1','n1','t1',NOW());
INSERT INTO public.news_tags(guid, guid_news, guid_tag, created_at) VALUES('nt2','n1','t2',NOW());
INSERT INTO public.news_tags(guid, guid_news, guid_tag, created_at) VALUES('nt3','n2','t1',NOW());

INSERT INTO public.topics(guid, name, created_at) VALUES('to1','Topic 1',NOW());

INSERT INTO public.topic_news(guid, guid_topic, guid_news, created_at) VALUES('tn1','to1','n1',NOW());
INSERT INTO public.topic_news(guid, guid_topic, guid_news, created_at) VALUES('tn2','to1','n2',NOW());