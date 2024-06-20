package geostore

import (
	"encoding/csv"
	"strings"
)

var statesCsvData = `
id,name,country_id,country_code,country_name,state_code,type,latitude,longitude
3901,Badakhshan,1,AF,Afghanistan,BDS,,36.73477250,70.81199530
3871,Badghis,1,AF,Afghanistan,BDG,,35.16713390,63.76953840
3875,Baghlan,1,AF,Afghanistan,BGL,,36.17890260,68.74530640
3884,Balkh,1,AF,Afghanistan,BAL,,36.75506030,66.89753720
3872,Bamyan,1,AF,Afghanistan,BAM,,34.81000670,67.82121040
3892,Daykundi,1,AF,Afghanistan,DAY,,33.66949500,66.04635340
3899,Farah,1,AF,Afghanistan,FRA,,32.49532800,62.26266270
3889,Faryab,1,AF,Afghanistan,FYB,,36.07956130,64.90595500
3870,Ghazni,1,AF,Afghanistan,GHA,,33.54505870,68.41739720
3888,Ghōr,1,AF,Afghanistan,GHO,,34.09957760,64.90595500
3873,Helmand,1,AF,Afghanistan,HEL,,39.29893610,-76.61604720
3887,Herat,1,AF,Afghanistan,HER,,34.35286500,62.20402870
3886,Jowzjan,1,AF,Afghanistan,JOW,,36.89696920,65.66585680
3902,Kabul,1,AF,Afghanistan,KAB,,34.55534940,69.20748600
3890,Kandahar,1,AF,Afghanistan,KAN,,31.62887100,65.73717490
3879,Kapisa,1,AF,Afghanistan,KAP,,34.98105720,69.62145620
3878,Khost,1,AF,Afghanistan,KHO,,33.33384720,69.93716730
3876,Kunar,1,AF,Afghanistan,KNR,,34.84658930,71.09731700
3900,"Kunduz Province",1,AF,Afghanistan,KDZ,,36.72855110,68.86789820
3891,Laghman,1,AF,Afghanistan,LAG,,34.68976870,70.14558050
3897,Logar,1,AF,Afghanistan,LOG,,34.01455180,69.19239160
3882,Nangarhar,1,AF,Afghanistan,NAN,,34.17183130,70.62167940
3896,Nimruz,1,AF,Afghanistan,NIM,,31.02614880,62.45041540
3880,Nuristan,1,AF,Afghanistan,NUR,,35.32502230,70.90712360
3894,Paktia,1,AF,Afghanistan,PIA,,33.70619900,69.38310790
3877,Paktika,1,AF,Afghanistan,PKA,,32.26453860,68.52471490
3881,Panjshir,1,AF,Afghanistan,PAN,,38.88023910,-77.17172380
3895,Parwan,1,AF,Afghanistan,PAR,,34.96309770,68.81088490
3883,Samangan,1,AF,Afghanistan,SAM,,36.31555060,67.96428630
3885,"Sar-e Pol",1,AF,Afghanistan,SAR,,36.21662800,65.93336000
3893,Takhar,1,AF,Afghanistan,TAK,,36.66980130,69.47845410
3898,Urozgan,1,AF,Afghanistan,URU,,32.92712870,66.14152630
3874,Zabul,1,AF,Afghanistan,ZAB,,32.19187820,67.18944880
603,"Berat County",3,AL,Albania,01,,40.69530120,20.04496620
629,"Berat District",3,AL,Albania,BR,,40.70863770,19.94373140
607,"Bulqizë District",3,AL,Albania,BU,,41.49425870,20.21471570
618,"Delvinë District",3,AL,Albania,DL,,39.94813640,20.09558910
608,"Devoll District",3,AL,Albania,DV,,40.64473470,20.95066360
610,"Dibër County",3,AL,Albania,09,,41.58881630,20.23556470
605,"Dibër District",3,AL,Albania,DI,,41.58881630,20.23556470
632,"Durrës County",3,AL,Albania,02,,41.50809720,19.61631850
639,"Durrës District",3,AL,Albania,DR,,41.37065170,19.52110630
598,"Elbasan County",3,AL,Albania,03,,41.12666720,20.23556470
631,"Fier County",3,AL,Albania,04,,40.91913920,19.66393090
627,"Fier District",3,AL,Albania,FR,,40.72750400,19.56275960
604,"Gjirokastër County",3,AL,Albania,05,,40.06728740,20.10452290
621,"Gjirokastër District",3,AL,Albania,GJ,,40.06728740,20.10452290
617,"Gramsh District",3,AL,Albania,GR,,40.86698730,20.18493230
600,"Has District",3,AL,Albania,HA,,42.79013360,-83.61220120
594,"Kavajë District",3,AL,Albania,KA,,41.18445290,19.56275960
628,"Kolonjë District",3,AL,Albania,ER,,40.33732620,20.67946760
630,"Korçë County",3,AL,Albania,06,,40.59056700,20.61689210
597,"Korçë District",3,AL,Albania,KO,,40.59056700,20.61689210
614,"Krujë District",3,AL,Albania,KR,,41.50947650,19.77107320
612,"Kuçovë District",3,AL,Albania,KC,,40.78370630,19.87823480
601,"Kukës County",3,AL,Albania,07,,42.08074640,20.41429230
623,"Kukës District",3,AL,Albania,KU,,42.08074640,20.41429230
622,"Kurbin District",3,AL,Albania,KB,,41.64126440,19.70559500
609,"Lezhë County",3,AL,Albania,08,,41.78137590,19.80679160
595,"Lezhë District",3,AL,Albania,LE,,41.78607300,19.64607580
596,"Librazhd District",3,AL,Albania,LB,,41.18292320,20.31747690
599,"Lushnjë District",3,AL,Albania,LU,,40.94198300,19.69964280
602,"Malësi e Madhe District",3,AL,Albania,MM,,42.42451730,19.61631850
637,"Mallakastër District",3,AL,Albania,MK,,40.52733760,19.78297910
635,"Mat District",3,AL,Albania,MT,,41.59376750,19.99732440
638,"Mirditë District",3,AL,Albania,MR,,41.76428600,19.90205090
619,"Peqin District",3,AL,Albania,PQ,,41.04709020,19.75023840
625,"Përmet District",3,AL,Albania,PR,,40.23618370,20.35173340
606,"Pogradec District",3,AL,Albania,PG,,40.90153140,20.65562890
620,"Pukë District",3,AL,Albania,PU,,42.04697720,19.89609680
624,"Sarandë District",3,AL,Albania,SR,,39.85921190,20.02710010
611,"Shkodër County",3,AL,Albania,10,,42.15037100,19.66393090
626,"Shkodër District",3,AL,Albania,SH,,42.06929850,19.50325590
593,"Skrapar District",3,AL,Albania,SK,,40.53499460,20.28322170
616,"Tepelenë District",3,AL,Albania,TE,,40.29666320,20.01816730
615,"Tirana County",3,AL,Albania,11,,41.24275980,19.80679160
633,"Tirana District",3,AL,Albania,TR,,41.32754590,19.81869820
636,"Tropojë District",3,AL,Albania,TP,,42.39821510,20.16259550
634,"Vlorë County",3,AL,Albania,12,,40.15009600,19.80679160
613,"Vlorë District",3,AL,Albania,VL,,40.46606680,19.49135600
1118,Adrar,4,DZ,Algeria,01,,26.41813100,-0.60147170
1119,"Aïn Defla",4,DZ,Algeria,44,,36.25094290,1.93938150
1122,"Aïn Témouchent",4,DZ,Algeria,46,,35.29926980,-1.13927920
1144,Algiers,4,DZ,Algeria,16,,36.69972940,3.05761990
1103,Annaba,4,DZ,Algeria,23,,36.80205080,7.52472430
1142,Batna,4,DZ,Algeria,05,,35.59659540,5.89871390
1108,Béchar,4,DZ,Algeria,08,,31.62380980,-2.21624430
1128,Béjaïa,4,DZ,Algeria,06,,36.75152580,5.05568370
4909,"Béni Abbès",4,DZ,Algeria,53,,30.08310420,-2.83450520
1114,Biskra,4,DZ,Algeria,07,,34.84494370,5.72485670
1111,Blida,4,DZ,Algeria,09,,36.53112300,2.89762540
4908,"Bordj Baji Mokhtar",4,DZ,Algeria,52,,22.96633500,-3.94727320
1116,"Bordj Bou Arréridj",4,DZ,Algeria,34,,36.07399250,4.76302710
1104,Bouïra,4,DZ,Algeria,10,,36.36918460,3.90061940
1125,Boumerdès,4,DZ,Algeria,35,,36.68395590,3.62178020
1105,Chlef,4,DZ,Algeria,02,,36.16935150,1.28910360
1121,Constantine,4,DZ,Algeria,25,,36.33739110,6.66381200
4912,Djanet,4,DZ,Algeria,56,,23.83108720,8.70046720
1098,Djelfa,4,DZ,Algeria,17,,34.67039560,3.25037610
1129,"El Bayadh",4,DZ,Algeria,32,,32.71488240,0.90566230
4905,"El M'ghair",4,DZ,Algeria,49,,33.95405610,5.13464180
4906,"El Menia",4,DZ,Algeria,50,,31.36422500,2.57844950
1099,"El Oued",4,DZ,Algeria,39,,33.36781100,6.85165110
1100,"El Tarf",4,DZ,Algeria,36,,36.75766780,8.30763430
1127,Ghardaïa,4,DZ,Algeria,47,,32.49437410,3.64446000
1137,Guelma,4,DZ,Algeria,24,,36.46274440,7.43308330
1112,Illizi,4,DZ,Algeria,33,,26.16900050,8.48424650
4914,"In Guezzam",4,DZ,Algeria,58,,20.38643230,4.77893940
4913,"In Salah",4,DZ,Algeria,57,,27.21492290,1.84843960
1113,Jijel,4,DZ,Algeria,18,,36.71796810,5.98325770
1126,Khenchela,4,DZ,Algeria,40,,35.42694040,7.14601550
1138,Laghouat,4,DZ,Algeria,03,,33.80783410,2.86282940
1134,M'Sila,4,DZ,Algeria,28,,35.71866460,4.52334230
1124,Mascara,4,DZ,Algeria,29,,35.39041250,0.14949880
1109,Médéa,4,DZ,Algeria,26,,36.26370780,2.75878570
1132,Mila,4,DZ,Algeria,43,,36.36479570,6.15269850
1140,Mostaganem,4,DZ,Algeria,27,,35.95830540,0.33718890
1102,Naama,4,DZ,Algeria,45,,33.26673170,-0.31286590
1101,Oran,4,DZ,Algeria,31,,35.60823510,-0.56360900
1139,Ouargla,4,DZ,Algeria,30,,32.22648630,5.72998210
4907,"Ouled Djellal",4,DZ,Algeria,51,,34.41782210,4.96858430
1136,"Oum El Bouaghi",4,DZ,Algeria,04,,35.86887890,7.11082660
1130,Relizane,4,DZ,Algeria,48,,35.73834050,0.75328090
1123,Saïda,4,DZ,Algeria,20,,34.84152070,0.14560550
1141,Sétif,4,DZ,Algeria,19,,36.30733890,5.56172790
4902,"Sidi Bel Abbès",4,DZ,Algeria,22,,34.68060240,-1.09994950
1110,Skikda,4,DZ,Algeria,21,,36.67211980,6.83509990
1143,"Souk Ahras",4,DZ,Algeria,41,,36.28010620,7.93840330
1135,Tamanghasset,4,DZ,Algeria,11,,22.79029720,5.51932680
1117,Tébessa,4,DZ,Algeria,12,,35.12906910,7.95928630
1106,Tiaret,4,DZ,Algeria,14,,35.37086890,1.32178520
4910,Timimoun,4,DZ,Algeria,54,,29.67890600,0.50046080
1120,Tindouf,4,DZ,Algeria,37,,27.80631190,-5.72998210
1115,Tipasa,4,DZ,Algeria,42,,36.54626500,2.18432850
1133,Tissemsilt,4,DZ,Algeria,38,,35.60537810,1.81309800
1131,"Tizi Ouzou",4,DZ,Algeria,15,,36.70691100,4.23333550
1107,Tlemcen,4,DZ,Algeria,13,,34.67802840,-1.36621600
4911,Touggourt,4,DZ,Algeria,55,,33.12484760,5.78327150
488,"Andorra la Vella",6,AD,Andorra,07,,42.50631740,1.52183550
489,Canillo,6,AD,Andorra,02,,42.59782490,1.65663770
487,Encamp,6,AD,Andorra,03,,42.53597640,1.58367730
492,Escaldes-Engordany,6,AD,Andorra,08,,42.49093790,1.58869660
493,"La Massana",6,AD,Andorra,04,,42.54562500,1.51473920
491,Ordino,6,AD,Andorra,05,,42.59944330,1.54023270
490,"Sant Julià de Lòria",6,AD,Andorra,06,,42.45296310,1.49182350
221,"Bengo Province",7,AO,Angola,BGO,,-9.10422570,13.72891670
218,"Benguela Province",7,AO,Angola,BGU,,-12.80037440,13.91439900
212,"Bié Province",7,AO,Angola,BIE,,-12.57279070,17.66888700
228,"Cabinda Province",7,AO,Angola,CAB,,-5.02487490,12.34638750
226,"Cuando Cubango Province",7,AO,Angola,CCU,,-16.41808240,18.80761950
217,"Cuanza Norte Province",7,AO,Angola,CNO,,-9.23985130,14.65878210
216,"Cuanza Sul",7,AO,Angola,CUS,,-10.59519100,15.40680790
215,"Cunene Province",7,AO,Angola,CNN,,-16.28022210,16.15809370
213,"Huambo Province",7,AO,Angola,HUA,,-12.52682210,15.59433880
225,"Huíla Province",7,AO,Angola,HUI,,-14.92805530,14.65878210
222,"Luanda Province",7,AO,Angola,LUA,,-9.03508800,13.26634790
223,"Lunda Norte Province",7,AO,Angola,LNO,,-8.35250220,19.18800470
220,"Lunda Sul Province",7,AO,Angola,LSU,,-10.28665780,20.71224650
227,"Malanje Province",7,AO,Angola,MAL,,-9.82511830,16.91225100
219,"Moxico Province",7,AO,Angola,MOX,,-13.42935790,20.33088140
224,"Uíge Province",7,AO,Angola,UIG,,-7.17367320,15.40680790
214,"Zaire Province",7,AO,Angola,ZAI,,-6.57334580,13.17403480
3708,Barbuda,10,AG,"Antigua and Barbuda",10,,17.62662420,-61.77130280
3703,Redonda,10,AG,"Antigua and Barbuda",11,,16.93841600,-62.34551480
3709,"Saint George Parish",10,AG,"Antigua and Barbuda",03,,,
3706,"Saint John Parish",10,AG,"Antigua and Barbuda",04,,,
3707,"Saint Mary Parish",10,AG,"Antigua and Barbuda",05,,,
3705,"Saint Paul Parish",10,AG,"Antigua and Barbuda",06,,,
3704,"Saint Peter Parish",10,AG,"Antigua and Barbuda",07,,,
3710,"Saint Philip Parish",10,AG,"Antigua and Barbuda",08,,40.43682580,-80.06855320
3656,"Buenos Aires",11,AR,Argentina,B,province,-37.20172850,-59.84106970
3647,Catamarca,11,AR,Argentina,K,province,-28.47158770,-65.78772090
3640,Chaco,11,AR,Argentina,H,province,-27.42571750,-59.02437840
3651,Chubut,11,AR,Argentina,U,province,-43.29342460,-65.11148180
4880,"Ciudad Autónoma de Buenos Aires",11,AR,Argentina,C,city,-34.60368440,-58.38155910
3642,Córdoba,11,AR,Argentina,X,province,-31.39928760,-64.26438420
3638,Corrientes,11,AR,Argentina,W,province,-27.46921310,-58.83063490
3654,"Entre Ríos",11,AR,Argentina,E,province,-31.77466540,-60.49564610
3652,Formosa,11,AR,Argentina,P,province,-26.18948040,-58.22428060
3645,Jujuy,11,AR,Argentina,Y,province,-24.18433970,-65.30217700
3655,"La Pampa",11,AR,Argentina,L,province,-36.61475730,-64.28392090
3653,"La Rioja",11,AR,Argentina,F,province,-29.41937930,-66.85599320
3646,Mendoza,11,AR,Argentina,M,province,-32.88945870,-68.84583860
3644,Misiones,11,AR,Argentina,N,province,-27.42692550,-55.94670760
3648,Neuquén,11,AR,Argentina,Q,province,-38.94587000,-68.07309250
3639,"Río Negro",11,AR,Argentina,R,province,-40.82614340,-63.02663390
3643,Salta,11,AR,Argentina,A,province,-24.79976880,-65.41503670
3634,"San Juan",11,AR,Argentina,J,province,-31.53169760,-68.56769620
3636,"San Luis",11,AR,Argentina,D,province,-33.29620420,-66.32949480
3649,"Santa Cruz",11,AR,Argentina,Z,province,-51.63528210,-69.24743530
3641,"Santa Fe",11,AR,Argentina,S,province,-31.58551090,-60.72380160
3635,"Santiago del Estero",11,AR,Argentina,G,province,-27.78335740,-64.26416700
3650,"Tierra del Fuego",11,AR,Argentina,V,province,-54.80539980,-68.32420610
3637,Tucumán,11,AR,Argentina,T,province,-26.82211270,-65.21929030
2023,"Aragatsotn Region",12,AM,Armenia,AG,,40.33473010,44.37482960
2024,"Ararat Province",12,AM,Armenia,AR,,39.91394150,44.72000040
2026,"Armavir Region",12,AM,Armenia,AV,,40.15546310,44.03724460
2028,"Gegharkunik Province",12,AM,Armenia,GR,,40.35264260,45.12604140
2033,"Kotayk Region",12,AM,Armenia,KT,,40.54102140,44.76901480
2029,"Lori Region",12,AM,Armenia,LO,,40.96984520,44.49001380
2031,"Shirak Region",12,AM,Armenia,SH,,40.96308140,43.81024610
2027,"Syunik Province",12,AM,Armenia,SU,,39.51331120,46.33932340
2032,"Tavush Region",12,AM,Armenia,TV,,40.88662960,45.33934900
2025,"Vayots Dzor Region",12,AM,Armenia,VD,,39.76419960,45.33375280
2030,Yerevan,12,AM,Armenia,ER,,40.18720230,44.51520900
3907,"Australian Capital Territory",14,AU,Australia,ACT,territory,-35.47346790,149.01236790
3909,"New South Wales",14,AU,Australia,NSW,state,-31.25321830,146.92109900
3910,"Northern Territory",14,AU,Australia,NT,territory,-19.49141080,132.55096030
3905,Queensland,14,AU,Australia,QLD,state,-20.91757380,142.70279560
3904,"South Australia",14,AU,Australia,SA,state,-30.00023150,136.20915470
3908,Tasmania,14,AU,Australia,TAS,state,-41.45451960,145.97066470
3903,Victoria,14,AU,Australia,VIC,state,-36.48564230,140.97794250
3906,"Western Australia",14,AU,Australia,WA,state,-27.67281680,121.62830980
2062,Burgenland,15,AT,Austria,1,,47.15371650,16.26887970
2057,Carinthia,15,AT,Austria,2,,46.72220300,14.18058820
2065,"Lower Austria",15,AT,Austria,3,,48.10807700,15.80495580
2061,Salzburg,15,AT,Austria,5,,47.80949000,13.05501000
2059,Styria,15,AT,Austria,6,,47.35934420,14.46998270
2064,Tyrol,15,AT,Austria,7,,47.25374140,11.60148700
2058,"Upper Austria",15,AT,Austria,4,,48.02585400,13.97236650
2060,Vienna,15,AT,Austria,9,,48.20817430,16.37381890
2063,Vorarlberg,15,AT,Austria,8,,47.24974270,9.97973730
540,"Absheron District",16,AZ,Azerbaijan,ABS,,40.36296930,49.27368150
559,"Agdam District",16,AZ,Azerbaijan,AGM,,39.99318530,46.99495620
553,"Agdash District",16,AZ,Azerbaijan,AGS,,40.63354270,47.46743100
577,"Aghjabadi District",16,AZ,Azerbaijan,AGC,,28.78918410,77.51607880
543,"Agstafa District",16,AZ,Azerbaijan,AGA,,41.26559330,45.51342910
547,"Agsu District",16,AZ,Azerbaijan,AGU,,40.52833390,48.36508350
528,"Astara District",16,AZ,Azerbaijan,AST,,38.49378450,48.69443650
575,"Babek District",16,AZ,Azerbaijan,BAB,,39.15076130,45.44853680
552,Baku,16,AZ,Azerbaijan,BA,,40.40926170,49.86709240
560,"Balakan District",16,AZ,Azerbaijan,BAL,,41.70375090,46.40442130
569,"Barda District",16,AZ,Azerbaijan,BAR,,40.37065550,47.13789090
554,"Beylagan District",16,AZ,Azerbaijan,BEY,,39.77230730,47.61541660
532,"Bilasuvar District",16,AZ,Azerbaijan,BIL,,39.45988330,48.55098130
561,"Dashkasan District",16,AZ,Azerbaijan,DAS,,40.52022570,46.07793040
527,"Fizuli District",16,AZ,Azerbaijan,FUZ,,39.53786050,47.30338770
585,Ganja,16,AZ,Azerbaijan,GA,,36.36873380,-95.99857670
589,Gədəbəy,16,AZ,Azerbaijan,GAD,,40.56996390,45.81068830
573,"Gobustan District",16,AZ,Azerbaijan,QOB,,40.53261040,48.92737500
551,"Goranboy District",16,AZ,Azerbaijan,GOR,,40.53805060,46.59908910
531,Goychay,16,AZ,Azerbaijan,GOY,,40.62361680,47.74030340
574,"Goygol District",16,AZ,Azerbaijan,GYG,,40.55953780,46.33149530
571,"Hajigabul District",16,AZ,Azerbaijan,HAC,,40.03937700,48.92025330
544,"Imishli District",16,AZ,Azerbaijan,IMI,,39.86946860,48.06652180
564,"Ismailli District",16,AZ,Azerbaijan,ISM,,40.74299360,48.21255560
570,"Jabrayil District",16,AZ,Azerbaijan,CAB,,39.26455440,46.96215620
578,"Jalilabad District",16,AZ,Azerbaijan,CAL,,39.20516320,48.51006040
572,"Julfa District",16,AZ,Azerbaijan,CUL,,38.96049830,45.62929390
525,"Kalbajar District",16,AZ,Azerbaijan,KAL,,40.10243290,46.03648720
567,"Kangarli District",16,AZ,Azerbaijan,KAN,,39.38719400,45.16398520
590,"Khachmaz District",16,AZ,Azerbaijan,XAC,,41.45911680,48.80206260
537,"Khizi District",16,AZ,Azerbaijan,XIZ,,40.91094890,49.07292640
524,"Khojali District",16,AZ,Azerbaijan,XCI,,39.91325530,46.79430500
549,"Kurdamir District",16,AZ,Azerbaijan,KUR,,40.36986510,48.16446260
541,"Lachin District",16,AZ,Azerbaijan,LAC,,39.63834140,46.54608530
587,Lankaran,16,AZ,Azerbaijan,LAN,,38.75286690,48.84750150
558,"Lankaran District",16,AZ,Azerbaijan,LA,,38.75286690,48.84750150
546,"Lerik District",16,AZ,Azerbaijan,LER,,38.77361920,48.41514830
568,Martuni,16,AZ,Azerbaijan,XVD,,39.79146930,47.11008140
555,"Masally District",16,AZ,Azerbaijan,MAS,,39.03407220,48.65893540
580,Mingachevir,16,AZ,Azerbaijan,MI,,40.77025630,47.04960150
562,"Nakhchivan Autonomous Republic",16,AZ,Azerbaijan,NX,,39.32568140,45.49126480
530,"Neftchala District",16,AZ,Azerbaijan,NEF,,39.38810520,49.24137430
556,"Oghuz District",16,AZ,Azerbaijan,OGU,,41.07279240,47.46506720
534,"Ordubad District",16,AZ,Azerbaijan,ORD,,38.90216220,46.02376250
542,"Qabala District",16,AZ,Azerbaijan,QAB,,40.92539250,47.80161060
526,"Qakh District",16,AZ,Azerbaijan,QAX,,41.42068270,46.93201840
521,"Qazakh District",16,AZ,Azerbaijan,QAZ,,41.09710740,45.35163310
563,"Quba District",16,AZ,Azerbaijan,QBA,,41.15642420,48.41350210
548,"Qubadli District",16,AZ,Azerbaijan,QBI,,39.27139960,46.63543120
588,"Qusar District",16,AZ,Azerbaijan,QUS,,41.42668860,48.43455770
557,"Saatly District",16,AZ,Azerbaijan,SAT,,39.90955030,48.35951220
565,"Sabirabad District",16,AZ,Azerbaijan,SAB,,39.98706630,48.46925450
522,"Sadarak District",16,AZ,Azerbaijan,SAD,,39.71051140,44.88642770
545,"Salyan District",16,AZ,Azerbaijan,SAL,,28.35248110,82.12784000
536,"Samukh District",16,AZ,Azerbaijan,SMX,,40.76046310,46.40631810
591,"Shabran District",16,AZ,Azerbaijan,SBN,,41.22283760,48.84573040
579,"Shahbuz District",16,AZ,Azerbaijan,SAH,,39.44521030,45.65680090
518,Shaki,16,AZ,Azerbaijan,SA,,41.19747530,47.15712410
586,"Shaki District",16,AZ,Azerbaijan,SAK,,41.11346620,47.13169270
529,"Shamakhi District",16,AZ,Azerbaijan,SMI,,40.63187310,48.63638010
583,"Shamkir District",16,AZ,Azerbaijan,SKR,,40.82881440,46.01668790
535,"Sharur District",16,AZ,Azerbaijan,SAR,,39.55363320,44.98456800
520,Shirvan,16,AZ,Azerbaijan,SR,,39.94697070,48.92239190
592,"Shusha District",16,AZ,Azerbaijan,SUS,,39.75374380,46.74647550
584,"Siazan District",16,AZ,Azerbaijan,SIY,,41.07838330,49.11184770
582,Sumqayit,16,AZ,Azerbaijan,SM,,40.58547650,49.63174110
519,"Tartar District",16,AZ,Azerbaijan,TAR,,40.34438750,46.93765190
533,"Tovuz District",16,AZ,Azerbaijan,TOV,,40.99545230,45.61659070
539,"Ujar District",16,AZ,Azerbaijan,UCA,,40.50675250,47.64896410
550,"Yardymli District",16,AZ,Azerbaijan,YAR,,38.90589170,48.24961270
538,Yevlakh,16,AZ,Azerbaijan,YE,,40.61966380,47.15003240
523,"Yevlakh District",16,AZ,Azerbaijan,YEV,,40.61966380,47.15003240
581,"Zangilan District",16,AZ,Azerbaijan,ZAN,,39.08568990,46.65247280
566,"Zaqatala District",16,AZ,Azerbaijan,ZAQ,,41.59068890,46.72403730
576,"Zardab District",16,AZ,Azerbaijan,ZAR,,40.21481140,47.71494400
1992,Capital,18,BH,Bahrain,13,,,
1996,Central,18,BH,Bahrain,16,,26.14260930,50.56532940
1995,Muharraq,18,BH,Bahrain,15,,26.26856530,50.64825170
1994,Northern,18,BH,Bahrain,17,,26.15519140,50.48251730
1993,Southern,18,BH,Bahrain,14,,25.93810180,50.57568870
796,"Bagerhat District",19,BD,Bangladesh,05,district,22.66024360,89.78954780
802,Bahadia,19,BD,Bangladesh,33,,23.78987120,90.16714830
752,"Bandarban District",19,BD,Bangladesh,01,district,21.83110020,92.36863210
784,"Barguna District",19,BD,Bangladesh,02,district,22.09529150,90.11206960
818,"Barisal District",19,BD,Bangladesh,06,district,22.70220980,90.36963160
807,"Barisal Division",19,BD,Bangladesh,A,division,22.38111310,90.33718890
756,"Bhola District",19,BD,Bangladesh,07,district,22.17853150,90.71010230
797,"Bogra District",19,BD,Bangladesh,03,district,24.85104020,89.36972250
810,"Brahmanbaria District",19,BD,Bangladesh,04,district,23.96081810,91.11150140
768,"Chandpur District",19,BD,Bangladesh,09,district,23.25131480,90.85178460
761,"Chapai Nawabganj District",19,BD,Bangladesh,45,district,24.74131110,88.29120690
785,"Chittagong District",19,BD,Bangladesh,10,district,22.51501050,91.75388170
803,"Chittagong Division",19,BD,Bangladesh,B,division,23.17931570,91.98815270
788,"Chuadanga District",19,BD,Bangladesh,12,district,23.61605120,88.82630060
763,"Comilla District",19,BD,Bangladesh,08,district,23.45756670,91.18089960
751,"Cox's Bazar District",19,BD,Bangladesh,11,district,21.56406260,92.02821290
771,"Dhaka District",19,BD,Bangladesh,13,district,23.81051400,90.33718890
760,"Dhaka Division",19,BD,Bangladesh,C,division,23.95357420,90.14949880
783,"Dinajpur District",19,BD,Bangladesh,14,district,25.62791230,88.63317580
762,"Faridpur District",19,BD,Bangladesh,15,district,23.54239190,89.63089210
816,"Feni District",19,BD,Bangladesh,16,district,22.94087840,91.40666460
795,"Gaibandha District",19,BD,Bangladesh,19,district,25.32969280,89.54296520
798,"Gazipur District",19,BD,Bangladesh,18,district,24.09581710,90.41251810
792,"Gopalganj District",19,BD,Bangladesh,17,district,26.48315840,84.43655000
805,"Habiganj District",19,BD,Bangladesh,20,district,24.47712360,91.45065650
808,"Jamalpur District",19,BD,Bangladesh,21,district,25.08309260,89.78532180
757,"Jessore District",19,BD,Bangladesh,22,district,23.16340140,89.21816640
778,"Jhalokati District",19,BD,Bangladesh,25,district,22.57208000,90.18696440
789,"Jhenaidah District",19,BD,Bangladesh,23,district,23.54498730,89.17260310
806,"Joypurhat District",19,BD,Bangladesh,24,district,25.09473490,89.09449370
786,"Khagrachari District",19,BD,Bangladesh,29,district,23.13217510,91.94902100
811,"Khulna District",19,BD,Bangladesh,27,district,22.67377350,89.39665810
775,"Khulna Division",19,BD,Bangladesh,D,division,22.80878160,89.24671910
779,"Kishoreganj District",19,BD,Bangladesh,26,district,24.42604570,90.98206680
793,"Kurigram District",19,BD,Bangladesh,28,district,25.80724140,89.62947460
774,"Kushtia District",19,BD,Bangladesh,30,district,23.89069950,89.10993680
819,"Lakshmipur District",19,BD,Bangladesh,31,district,22.94467440,90.82819070
780,"Lalmonirhat District",19,BD,Bangladesh,32,district,25.99233980,89.28472510
817,"Madaripur District",19,BD,Bangladesh,36,district,23.23933460,90.18696440
776,"Meherpur District",19,BD,Bangladesh,39,district,23.80519910,88.67235780
794,"Moulvibazar District",19,BD,Bangladesh,38,district,24.30953440,91.73149030
790,"Munshiganj District",19,BD,Bangladesh,35,district,23.49809310,90.41266210
766,"Mymensingh District",19,BD,Bangladesh,34,district,24.75385750,90.40729190
758,"Mymensingh Division",19,BD,Bangladesh,H,division,24.71362000,90.45023680
814,"Naogaon District",19,BD,Bangladesh,48,district,24.91315970,88.75309520
769,"Narail District",19,BD,Bangladesh,43,district,23.11629290,89.58404040
770,"Narayanganj District",19,BD,Bangladesh,40,district,23.71466010,90.56360900
787,"Natore District",19,BD,Bangladesh,44,district,24.41024300,89.00761770
764,"Netrokona District",19,BD,Bangladesh,41,district,24.81032840,90.86564150
772,"Nilphamari District",19,BD,Bangladesh,46,district,25.84827980,88.94141340
815,"Noakhali District",19,BD,Bangladesh,47,district,22.87237890,91.09731840
754,"Pabna District",19,BD,Bangladesh,49,district,24.15850500,89.44807180
800,"Panchagarh District",19,BD,Bangladesh,52,district,26.27087050,88.59517510
777,"Patuakhali District",19,BD,Bangladesh,51,district,22.22486320,90.45475030
791,"Pirojpur District",19,BD,Bangladesh,50,district,22.57907440,89.97592640
773,"Rajbari District",19,BD,Bangladesh,53,district,23.71513400,89.58748190
813,"Rajshahi District",19,BD,Bangladesh,54,district,24.37330870,88.60487160
753,"Rajshahi Division",19,BD,Bangladesh,E,division,24.71057760,88.94138650
809,"Rangamati Hill District",19,BD,Bangladesh,56,district,22.73241730,92.29851340
759,"Rangpur District",19,BD,Bangladesh,55,district,25.74679250,89.25083350
750,"Rangpur Division",19,BD,Bangladesh,F,division,25.84833880,88.94138650
799,"Satkhira District",19,BD,Bangladesh,58,district,22.31548120,89.11145250
801,"Shariatpur District",19,BD,Bangladesh,62,district,23.24232140,90.43477110
755,"Sherpur District",19,BD,Bangladesh,57,district,25.07462350,90.14949040
781,"Sirajganj District",19,BD,Bangladesh,59,district,24.31411150,89.56996150
812,"Sunamganj District",19,BD,Bangladesh,61,district,25.07145350,91.39916270
767,"Sylhet District",19,BD,Bangladesh,60,district,24.89933570,91.87004730
765,"Sylhet Division",19,BD,Bangladesh,G,division,24.70498110,91.67606910
782,"Tangail District",19,BD,Bangladesh,63,district,24.39174270,89.99482570
804,"Thakurgaon District",19,BD,Bangladesh,64,district,26.04183920,88.42826160
1228,"Christ Church",20,BB,Barbados,01,,36.00604070,-95.92112100
1229,"Saint Andrew",20,BB,Barbados,02,,,
1226,"Saint George",20,BB,Barbados,03,,37.09652780,-113.56841640
1224,"Saint James",20,BB,Barbados,04,,48.52356600,-1.32378850
1227,"Saint John",20,BB,Barbados,05,,45.27331530,-66.06330800
1223,"Saint Joseph",20,BB,Barbados,06,,39.76745780,-94.84668100
1221,"Saint Lucy",20,BB,Barbados,07,,38.76146250,-77.44914390
1230,"Saint Michael",20,BB,Barbados,08,,36.03597700,-95.84905200
1222,"Saint Peter",20,BB,Barbados,09,,37.08271190,-94.51712500
1220,"Saint Philip",20,BB,Barbados,10,,35.23311400,-89.43640420
1225,"Saint Thomas",20,BB,Barbados,11,,18.33809650,-64.89409460
2959,"Brest Region",21,BY,Belarus,BR,,52.52966410,25.46064800
2955,"Gomel Region",21,BY,Belarus,HO,,52.16487540,29.13332510
2956,"Grodno Region",21,BY,Belarus,HR,,53.65999450,25.34485710
2958,Minsk,21,BY,Belarus,HM,,53.90060110,27.55897200
2957,"Minsk Region",21,BY,Belarus,MI,,54.10678890,27.41292450
2954,"Mogilev Region",21,BY,Belarus,MA,,53.51017910,30.40064440
2960,"Vitebsk Region",21,BY,Belarus,VI,,55.29598330,28.75836270
1381,Antwerp,22,BE,Belgium,VAN,,51.21944750,4.40246430
1376,"Brussels-Capital Region",22,BE,Belgium,BRU,,50.85034630,4.35172110
1377,"East Flanders",22,BE,Belgium,VOV,,51.03621010,3.73731240
1373,Flanders,22,BE,Belgium,VLG,,51.01087060,3.72646130
1374,"Flemish Brabant",22,BE,Belgium,VBR,,50.88154340,4.56459700
1375,Hainaut,22,BE,Belgium,WHT,,50.52570760,4.06210170
1384,Liège,22,BE,Belgium,WLG,,50.63255740,5.57966620
1372,Limburg,22,BE,Belgium,VLI,,,
1379,Luxembourg,22,BE,Belgium,WLX,,49.81527300,6.12958300
1378,Namur,22,BE,Belgium,WNA,,50.46738830,4.87198540
1380,Wallonia,22,BE,Belgium,WAL,,50.41756370,4.45100660
1382,"Walloon Brabant",22,BE,Belgium,WBR,,50.63324100,4.52431500
1383,"West Flanders",22,BE,Belgium,VWV,,51.04047470,2.99942130
264,"Belize District",23,BZ,Belize,BZ,,17.56776790,-88.40160410
269,"Cayo District",23,BZ,Belize,CY,,17.09844450,-88.94138650
266,"Corozal District",23,BZ,Belize,CZL,,18.13492380,-88.24611830
268,"Orange Walk District",23,BZ,Belize,OW,,17.76035300,-88.86469800
265,"Stann Creek District",23,BZ,Belize,SC,,16.81166310,-88.40160410
267,"Toledo District",23,BZ,Belize,TOL,,16.24919230,-88.86469800
3077,"Alibori Department",24,BJ,Benin,AL,,10.96810930,2.77798130
3076,"Atakora Department",24,BJ,Benin,AK,,10.79549310,1.67606910
3079,"Atlantique Department",24,BJ,Benin,AQ,,6.65883910,2.22366670
3078,"Borgou Department",24,BJ,Benin,BO,,9.53408640,2.77798130
3070,"Collines Department",24,BJ,Benin,CO,,8.30222970,2.30244600
3072,"Donga Department",24,BJ,Benin,DO,,9.71918670,1.67606910
3071,"Kouffo Department",24,BJ,Benin,KO,,7.00358940,1.75388170
3081,"Littoral Department",24,BJ,Benin,LI,,6.38069730,2.44063870
3075,"Mono Department",24,BJ,Benin,MO,,37.92186080,-118.95286450
3080,"Ouémé Department",24,BJ,Benin,OU,,6.61481520,2.49999180
3074,"Plateau Department",24,BJ,Benin,PL,,7.34451410,2.53960300
3073,"Zou Department",24,BJ,Benin,ZO,,7.34692680,2.06651970
4860,Devonshire,25,BM,Bermuda,DEV,municipality,32.30380620,-64.76069540
4861,Hamilton,25,BM,Bermuda,HA,municipality,32.34494320,-64.72365000
4863,Paget,25,BM,Bermuda,PAG,municipality,32.28107400,-64.77847870
4864,Pembroke,25,BM,Bermuda,PEM,municipality,32.30076720,-64.79626300
4866,"Saint George's",25,BM,Bermuda,SGE,municipality,17.12577590,-62.56198110
4867,Sandys,25,BM,Bermuda,SAN,municipality,32.29995280,-64.86741030
4868,Smith's,25,BM,Bermuda,SMI,municipality,32.31339660,-64.73105880
4869,Southampton,25,BM,Bermuda,SOU,municipality,32.25400950,-64.82590580
4870,Warwick,25,BM,Bermuda,WAR,municipality,32.26615340,-64.80811980
240,"Bumthang District",26,BT,Bhutan,33,,27.64183900,90.67730460
239,"Chukha District",26,BT,Bhutan,12,,27.07843040,89.47421770
238,"Dagana District",26,BT,Bhutan,22,,27.03228610,89.88793040
229,"Gasa District",26,BT,Bhutan,GA,,28.01858860,89.92532330
232,"Haa District",26,BT,Bhutan,13,,27.26516690,89.17059980
234,"Lhuntse District",26,BT,Bhutan,44,,27.82649890,91.13530200
242,"Mongar District",26,BT,Bhutan,42,,27.26170590,91.28910360
237,"Paro District",26,BT,Bhutan,11,,27.42859490,89.41665160
244,"Pemagatshel District",26,BT,Bhutan,43,,27.00238200,91.34692470
235,"Punakha District",26,BT,Bhutan,23,,27.69037160,89.88793040
243,"Samdrup Jongkhar District",26,BT,Bhutan,45,,26.80356820,91.50392070
246,"Samtse District",26,BT,Bhutan,14,,27.02918320,89.05615320
247,"Sarpang District",26,BT,Bhutan,31,,26.93730410,90.48799160
241,"Thimphu District",26,BT,Bhutan,15,,27.47122160,89.63390410
236,"Trashigang District",26,BT,Bhutan,41,,27.25667950,91.75388170
245,"Trongsa District",26,BT,Bhutan,32,,27.50022690,90.50806340
230,"Tsirang District",26,BT,Bhutan,21,,27.03220700,90.18696440
231,"Wangdue Phodrang District",26,BT,Bhutan,24,,27.45260460,90.06749280
233,"Zhemgang District",26,BT,Bhutan,34,,27.07697500,90.82940020
3375,"Beni Department",27,BO,Bolivia,B,,-14.37827470,-65.09577920
3382,"Chuquisaca Department",27,BO,Bolivia,H,,-20.02491440,-64.14782360
3381,"Cochabamba Department",27,BO,Bolivia,C,,-17.56816750,-65.47573600
3380,"La Paz Department",27,BO,Bolivia,L,,,
3376,"Oruro Department",27,BO,Bolivia,O,,-18.57115790,-67.76159830
3379,"Pando Department",27,BO,Bolivia,N,,-10.79889010,-66.99880110
3383,"Potosí Department",27,BO,Bolivia,P,,-20.62471300,-66.99880110
3377,"Santa Cruz Department",27,BO,Bolivia,S,,-16.74760370,-62.07509980
3378,"Tarija Department",27,BO,Bolivia,T,,-21.58315950,-63.95861110
5086,Bonaire,155,BQ,"Bonaire, Sint Eustatius and Saba",BQ1,"special municipality",12.20189020,-68.26238220
5087,Saba,155,BQ,"Bonaire, Sint Eustatius and Saba",BQ2,"special municipality",17.63546420,-63.23267630
5088,"Sint Eustatius",155,BQ,"Bonaire, Sint Eustatius and Saba",BQ3,"special municipality",17.48903060,-62.97355500
472,"Bosnian Podrinje Canton",28,BA,"Bosnia and Herzegovina",05,,43.68749000,18.82443940
460,"Brčko District",28,BA,"Bosnia and Herzegovina",BRC,,44.84059440,18.74215300
471,"Canton 10",28,BA,"Bosnia and Herzegovina",10,,43.95341550,16.94251870
462,"Central Bosnia Canton",28,BA,"Bosnia and Herzegovina",06,,44.13818560,17.68667140
467,"Federation of Bosnia and Herzegovina",28,BA,"Bosnia and Herzegovina",BIH,,43.88748970,17.84279300
463,"Herzegovina-Neretva Canton",28,BA,"Bosnia and Herzegovina",07,,43.52651590,17.76362100
464,"Posavina Canton",28,BA,"Bosnia and Herzegovina",02,,45.07520940,18.37763040
470,"Republika Srpska",28,BA,"Bosnia and Herzegovina",SRP,,44.72801860,17.31481360
466,"Sarajevo Canton",28,BA,"Bosnia and Herzegovina",09,,43.85125640,18.29534420
461,"Tuzla Canton",28,BA,"Bosnia and Herzegovina",03,,44.53434630,18.69727970
465,"Una-Sana Canton",28,BA,"Bosnia and Herzegovina",01,,44.65031160,16.31716290
469,"West Herzegovina Canton",28,BA,"Bosnia and Herzegovina",08,,43.43692440,17.38488310
468,"Zenica-Doboj Canton",28,BA,"Bosnia and Herzegovina",04,,44.21271090,18.16046250
3067,"Central District",29,BW,Botswana,CE,,,
3061,"Ghanzi District",29,BW,Botswana,GH,,-21.86523140,21.85685860
3066,"Kgalagadi District",29,BW,Botswana,KG,,-24.75502850,21.85685860
3062,"Kgatleng District",29,BW,Botswana,KL,,-24.19704450,26.23046160
3069,"Kweneng District",29,BW,Botswana,KW,,-23.83672490,25.28375850
3060,Ngamiland,29,BW,Botswana,NG,,-19.19053210,23.00119890
3068,"North-East District",29,BW,Botswana,NE,,37.58844610,-94.68637820
3065,"North-West District",29,BW,Botswana,NW,,39.34463070,-76.68542830
3064,"South-East District",29,BW,Botswana,SE,,31.21637980,-82.35270440
3063,"Southern District",29,BW,Botswana,SO,,,
2012,Acre,31,BR,Brazil,AC,,-9.02379600,-70.81199500
2007,Alagoas,31,BR,Brazil,AL,,-9.57130580,-36.78195050
1999,Amapá,31,BR,Brazil,AP,,0.90199250,-52.00295650
2004,Amazonas,31,BR,Brazil,AM,,-3.07000000,-61.66000000
2002,Bahia,31,BR,Brazil,BA,,-11.40987400,-41.28085700
2016,Ceará,31,BR,Brazil,CE,,-5.49839770,-39.32062410
2017,"Distrito Federal",31,BR,Brazil,DF,,-15.79976540,-47.86447150
2018,"Espírito Santo",31,BR,Brazil,ES,,-19.18342290,-40.30886260
2000,Goiás,31,BR,Brazil,GO,,-15.82703690,-49.83622370
2015,Maranhão,31,BR,Brazil,MA,,-4.96094980,-45.27441590
2011,"Mato Grosso",31,BR,Brazil,MT,,-12.68187120,-56.92109900
2010,"Mato Grosso do Sul",31,BR,Brazil,MS,,-20.77222950,-54.78515310
1998,"Minas Gerais",31,BR,Brazil,MG,,-18.51217800,-44.55503080
2009,Pará,31,BR,Brazil,PA,,-1.99812710,-54.93061520
2005,Paraíba,31,BR,Brazil,PB,,-7.23996090,-36.78195050
2022,Paraná,31,BR,Brazil,PR,,-25.25208880,-52.02154150
2006,Pernambuco,31,BR,Brazil,PE,,-8.81371730,-36.95410700
2008,Piauí,31,BR,Brazil,PI,,-7.71834010,-42.72892360
1997,"Rio de Janeiro",31,BR,Brazil,RJ,,-22.90684670,-43.17289650
2019,"Rio Grande do Norte",31,BR,Brazil,RN,,-5.40258030,-36.95410700
2001,"Rio Grande do Sul",31,BR,Brazil,RS,,-30.03463160,-51.21769860
2013,Rondônia,31,BR,Brazil,RO,,-11.50573410,-63.58061100
4858,Roraima,31,BR,Brazil,RR,,2.73759710,-62.07509980
2014,"Santa Catarina",31,BR,Brazil,SC,,-27.33000000,-49.44000000
2021,"São Paulo",31,BR,Brazil,SP,,-23.55051990,-46.63330940
2003,Sergipe,31,BR,Brazil,SE,,-10.57409340,-37.38565810
2020,Tocantins,31,BR,Brazil,TO,,-10.17528000,-48.29824740
1217,"Belait District",33,BN,Brunei,BE,,4.37507490,114.61928990
1216,"Brunei-Muara District",33,BN,Brunei,BM,,4.93112060,114.95168690
1218,"Temburong District",33,BN,Brunei,TE,,4.62041280,115.14148400
1219,"Tutong District",33,BN,Brunei,TU,,4.71403730,114.66679390
4699,"Blagoevgrad Province",34,BG,Bulgaria,01,,42.02086140,23.09433560
4715,"Burgas Province",34,BG,Bulgaria,02,,42.50480000,27.46260790
4718,"Dobrich Province",34,BG,Bulgaria,08,,43.57278600,27.82728020
4693,"Gabrovo Province",34,BG,Bulgaria,07,,42.86847000,25.31688900
4704,"Haskovo Province",34,BG,Bulgaria,26,,41.93441780,25.55546720
4702,"Kardzhali Province",34,BG,Bulgaria,09,,41.63384160,25.37766870
4703,"Kyustendil Province",34,BG,Bulgaria,10,,42.28687990,22.69396350
4710,"Lovech Province",34,BG,Bulgaria,11,,43.13677980,24.71393350
4696,"Montana Province",34,BG,Bulgaria,12,,43.40851480,23.22575890
4712,"Pazardzhik Province",34,BG,Bulgaria,13,,42.19275670,24.33362260
4695,"Pernik Province",34,BG,Bulgaria,14,,42.60519900,23.03779160
4706,"Pleven Province",34,BG,Bulgaria,15,,43.41701690,24.60667080
4701,"Plovdiv Province",34,BG,Bulgaria,16,,42.13540790,24.74529040
4698,"Razgrad Province",34,BG,Bulgaria,17,,43.52717050,26.52412280
4713,"Ruse Province",34,BG,Bulgaria,18,,43.83559640,25.96561440
4882,Shumen,34,BG,Bulgaria,27,,43.27123980,26.93612860
4708,"Silistra Province",34,BG,Bulgaria,19,,44.11471010,27.26714540
4700,"Sliven Province",34,BG,Bulgaria,20,,42.68167020,26.32285690
4694,"Smolyan Province",34,BG,Bulgaria,21,,41.57741480,24.70108710
4705,"Sofia City Province",34,BG,Bulgaria,22,,42.75701090,23.45046830
4719,"Sofia Province",34,BG,Bulgaria,23,,42.67344000,23.83349370
4707,"Stara Zagora Province",34,BG,Bulgaria,24,,42.42577090,25.63448550
4714,"Targovishte Province",34,BG,Bulgaria,25,,43.24623490,26.56912510
4717,"Varna Province",34,BG,Bulgaria,03,,43.20464770,27.91054880
4709,"Veliko Tarnovo Province",34,BG,Bulgaria,04,,43.07565390,25.61715000
4697,"Vidin Province",34,BG,Bulgaria,05,,43.99617390,22.86795150
4711,"Vratsa Province",34,BG,Bulgaria,06,,43.21018060,23.55292100
4716,"Yambol Province",34,BG,Bulgaria,28,,42.48414940,26.50352960
3160,"Balé Province",35,BF,"Burkina Faso",BAL,,11.78206020,-3.01757120
3155,"Bam Province",35,BF,"Burkina Faso",BAM,,13.44613300,-1.59839590
3120,"Banwa Province",35,BF,"Burkina Faso",BAN,,12.13230530,-4.15137640
3152,"Bazèga Province",35,BF,"Burkina Faso",BAZ,,11.97676920,-1.44346900
3138,"Boucle du Mouhoun Region",35,BF,"Burkina Faso",01,,12.41660000,-3.41955270
3121,"Bougouriba Province",35,BF,"Burkina Faso",BGR,,10.87226460,-3.33889170
3131,Boulgou,35,BF,"Burkina Faso",BLG,,11.43367660,-0.37483540
3153,"Cascades Region",35,BF,"Burkina Faso",02,,10.40729920,-4.56244260
3136,Centre,35,BF,"Burkina Faso",03,,,
3162,"Centre-Est Region",35,BF,"Burkina Faso",04,,11.52476740,-0.14949880
3127,"Centre-Nord Region",35,BF,"Burkina Faso",05,,13.17244640,-0.90566230
3115,"Centre-Ouest Region",35,BF,"Burkina Faso",06,,11.87984660,-2.30244600
3149,"Centre-Sud Region",35,BF,"Burkina Faso",07,,11.52289110,-1.05861350
3167,"Comoé Province",35,BF,"Burkina Faso",COM,,10.40729920,-4.56244260
3158,"Est Region",35,BF,"Burkina Faso",08,,12.43655260,0.90566230
3148,"Ganzourgou Province",35,BF,"Burkina Faso",GAN,,12.25376480,-0.75328090
3122,"Gnagna Province",35,BF,"Burkina Faso",GNA,,12.89749920,0.07467670
3143,"Gourma Province",35,BF,"Burkina Faso",GOU,,12.16244730,0.67730460
3165,"Hauts-Bassins Region",35,BF,"Burkina Faso",09,,11.49420030,-4.23333550
3129,"Houet Province",35,BF,"Burkina Faso",HOU,,11.13204470,-4.23333550
3135,"Ioba Province",35,BF,"Burkina Faso",IOB,,11.05620340,-3.01757120
3168,"Kadiogo Province",35,BF,"Burkina Faso",KAD,,12.34258970,-1.44346900
3112,"Kénédougou Province",35,BF,"Burkina Faso",KEN,,11.39193950,-4.97665400
3132,"Komondjari Province",35,BF,"Burkina Faso",KMD,,12.71265270,0.67730460
3157,"Kompienga Province",35,BF,"Burkina Faso",KMP,,11.52383620,0.75328090
3146,"Kossi Province",35,BF,"Burkina Faso",KOS,,12.96045800,-3.90626880
3133,"Koulpélogo Province",35,BF,"Burkina Faso",KOP,,11.52476740,0.14949880
3161,"Kouritenga Province",35,BF,"Burkina Faso",KOT,,12.16318130,-0.22446620
3147,"Kourwéogo Province",35,BF,"Burkina Faso",KOW,,12.70774950,-1.75388170
3159,"Léraba Province",35,BF,"Burkina Faso",LER,,10.66487850,-5.31025050
3151,"Loroum Province",35,BF,"Burkina Faso",LOR,,13.81298140,-2.06651970
3123,Mouhoun,35,BF,"Burkina Faso",MOU,,12.14323810,-3.33889170
3116,"Nahouri Province",35,BF,"Burkina Faso",NAO,,11.25022670,-1.13530200
3113,"Namentenga Province",35,BF,"Burkina Faso",NAM,,13.08125840,-0.52578230
3142,"Nayala Province",35,BF,"Burkina Faso",NAY,,12.69645580,-3.01757120
3164,"Nord Region, Burkina Faso",35,BF,"Burkina Faso",10,,13.71825200,-2.30244600
3156,"Noumbiel Province",35,BF,"Burkina Faso",NOU,,9.84409460,-2.97755580
3141,"Oubritenga Province",35,BF,"Burkina Faso",OUB,,12.70960870,-1.44346900
3144,"Oudalan Province",35,BF,"Burkina Faso",OUD,,14.47190200,-0.45023680
3117,"Passoré Province",35,BF,"Burkina Faso",PAS,,12.88812210,-2.22366670
3125,"Plateau-Central Region",35,BF,"Burkina Faso",11,,12.25376480,-0.75328090
3163,"Poni Province",35,BF,"Burkina Faso",PON,,10.33259960,-3.33889170
3114,"Sahel Region",35,BF,"Burkina Faso",12,,14.10008650,-0.14949880
3154,"Sanguié Province",35,BF,"Burkina Faso",SNG,,12.15018610,-2.69838680
3126,"Sanmatenga Province",35,BF,"Burkina Faso",SMT,,13.35653040,-1.05861350
3139,"Séno Province",35,BF,"Burkina Faso",SEN,,14.00722340,-0.07467670
3119,"Sissili Province",35,BF,"Burkina Faso",SIS,,11.24412190,-2.22366670
3166,"Soum Province",35,BF,"Burkina Faso",SOM,,14.09628410,-1.36621600
3137,"Sourou Province",35,BF,"Burkina Faso",SOR,,13.34180300,-2.93757390
3140,"Sud-Ouest Region",35,BF,"Burkina Faso",13,,10.42314930,-3.25836260
3128,"Tapoa Province",35,BF,"Burkina Faso",TAP,,12.24970720,1.67606910
3134,"Tuy Province",35,BF,"Burkina Faso",TUI,,38.88868400,-77.00471900
3124,"Yagha Province",35,BF,"Burkina Faso",YAG,,13.35761570,0.75328090
3150,"Yatenga Province",35,BF,"Burkina Faso",YAT,,13.62493440,-2.38136210
3145,"Ziro Province",35,BF,"Burkina Faso",ZIR,,11.60949950,-1.90992380
3130,"Zondoma Province",35,BF,"Burkina Faso",ZON,,13.11659260,-2.42087130
3118,"Zoundwéogo Province",35,BF,"Burkina Faso",ZOU,,11.61411740,-0.98206680
3196,"Bubanza Province",36,BI,Burundi,BB,,-3.15724030,29.37149090
3198,"Bujumbura Mairie Province",36,BI,Burundi,BM,,-3.38841410,29.34826460
3200,"Bujumbura Rural Province",36,BI,Burundi,BL,,-3.50901440,29.46435900
3202,"Bururi Province",36,BI,Burundi,BR,,-3.90068510,29.51077080
3201,"Cankuzo Province",36,BI,Burundi,CA,,-3.15277880,30.61998950
3190,"Cibitoke Province",36,BI,Burundi,CI,,-2.81028970,29.18557850
3197,"Gitega Province",36,BI,Burundi,GI,,-3.49290510,29.92779470
3194,"Karuzi Province",36,BI,Burundi,KR,,-3.13403470,30.11273500
3192,"Kayanza Province",36,BI,Burundi,KY,,-3.00779810,29.64991620
3195,"Kirundo Province",36,BI,Burundi,KI,,-2.57628820,30.11273500
3188,"Makamba Province",36,BI,Burundi,MA,,-4.32570620,29.69626770
3193,"Muramvya Province",36,BI,Burundi,MU,,-3.28983980,29.64991620
3186,"Muyinga Province",36,BI,Burundi,MY,,-2.77935110,30.29741990
3187,"Mwaro Province",36,BI,Burundi,MW,,-3.50259180,29.64991620
3199,"Ngozi Province",36,BI,Burundi,NG,,-2.89582430,29.88152030
3185,"Rumonge Province",36,BI,Burundi,RM,,-3.97540490,29.43880140
3189,"Rutana Province",36,BI,Burundi,RT,,-3.87915230,30.06652360
3191,"Ruyigi Province",36,BI,Burundi,RY,,-3.44620700,30.25127280
3984,"Banteay Meanchey",37,KH,Cambodia,1,province,13.75319140,102.98961500
3976,Battambang,37,KH,Cambodia,2,province,13.02869710,102.98961500
3991,"Kampong Cham",37,KH,Cambodia,3,province,12.09829180,105.31311850
3979,"Kampong Chhnang",37,KH,Cambodia,4,province,12.13923520,104.56552730
3988,"Kampong Speu",37,KH,Cambodia,5,province,11.61551090,104.37919120
5070,"Kampong Thom",37,KH,Cambodia,6,province,12.81674850,103.84131040
3981,Kampot,37,KH,Cambodia,7,province,10.73253510,104.37919120
3983,Kandal,37,KH,Cambodia,8,province,11.22373830,105.12589550
3978,Kep,37,KH,Cambodia,23,province,10.53608900,104.35591580
3982,"Koh Kong",37,KH,Cambodia,9,province,11.57628040,103.35872880
3986,Kratie,37,KH,Cambodia,10,province,12.50436080,105.96998780
3985,Mondulkiri,37,KH,Cambodia,11,province,12.78794270,107.10119310
3987,"Oddar Meanchey",37,KH,Cambodia,22,province,14.16097380,103.82162610
3980,Pailin,37,KH,Cambodia,24,province,12.90929620,102.66755750
3994,"Phnom Penh",37,KH,Cambodia,12,"autonomous municipality",11.55637380,104.92820990
3973,"Preah Vihear",37,KH,Cambodia,13,province,14.00857970,104.84546190
3974,"Prey Veng",37,KH,Cambodia,14,province,11.38024420,105.50054830
3977,Pursat,37,KH,Cambodia,15,province,12.27209560,103.72891670
3990,Ratanakiri,37,KH,Cambodia,16,province,13.85766070,107.10119310
3992,"Siem Reap",37,KH,Cambodia,17,province,13.33026600,104.10013260
3989,Sihanoukville,37,KH,Cambodia,18,province,10.75818990,103.82162610
3993,"Stung Treng",37,KH,Cambodia,19,province,13.57647300,105.96998780
3972,"Svay Rieng",37,KH,Cambodia,20,province,11.14272200,105.82902980
3975,Takeo,37,KH,Cambodia,21,province,10.93215190,104.79877100
2663,Adamawa,38,CM,Cameroon,AD,,9.32647510,12.39838530
2660,Centre,38,CM,Cameroon,CE,,,
2661,East,38,CM,Cameroon,ES,,39.01853360,-94.27924110
2656,"Far North",38,CM,Cameroon,EN,,66.76134510,124.12375300
2662,Littoral,38,CM,Cameroon,LT,,48.46227570,-68.51780710
2665,North,38,CM,Cameroon,NO,,37.09024000,-95.71289100
2657,Northwest,38,CM,Cameroon,NW,,36.37118570,-94.19346060
2659,South,38,CM,Cameroon,SU,,37.63159500,-97.34584090
2658,Southwest,38,CM,Cameroon,SW,,36.19088130,-95.88974480
2664,West,38,CM,Cameroon,OU,,37.03649890,-95.67059870
872,Alberta,39,CA,Canada,AB,province,53.93327060,-116.57650350
875,"British Columbia",39,CA,Canada,BC,province,53.72666830,-127.64762050
867,Manitoba,39,CA,Canada,MB,province,53.76086080,-98.81387620
868,"New Brunswick",39,CA,Canada,NB,province,46.56531630,-66.46191640
877,"Newfoundland and Labrador",39,CA,Canada,NL,province,53.13550910,-57.66043640
878,"Northwest Territories",39,CA,Canada,NT,territory,64.82554410,-124.84573340
874,"Nova Scotia",39,CA,Canada,NS,province,44.68198660,-63.74431100
876,Nunavut,39,CA,Canada,NU,territory,70.29977110,-83.10757700
866,Ontario,39,CA,Canada,ON,province,51.25377500,-85.32321400
871,"Prince Edward Island",39,CA,Canada,PE,province,46.51071200,-63.41681360
873,Quebec,39,CA,Canada,QC,province,52.93991590,-73.54913610
870,Saskatchewan,39,CA,Canada,SK,province,52.93991590,-106.45086390
869,Yukon,39,CA,Canada,YT,territory,35.50672150,-97.76254410
2994,"Barlavento Islands",40,CV,"Cape Verde",B,,16.82368450,-23.99348810
2999,"Boa Vista",40,CV,"Cape Verde",BV,,38.74346600,-120.73042970
2996,Brava,40,CV,"Cape Verde",BR,,40.98977780,-73.68357150
2991,"Maio Municipality",40,CV,"Cape Verde",MA,,15.20030980,-23.16797930
2987,Mosteiros,40,CV,"Cape Verde",MO,,37.89043480,-25.82075560
2997,Paul,40,CV,"Cape Verde",PA,,37.06250000,-95.67706800
2989,"Porto Novo",40,CV,"Cape Verde",PN,,6.49685740,2.62885230
2988,Praia,40,CV,"Cape Verde",PR,,14.93305000,-23.51332670
2982,"Ribeira Brava Municipality",40,CV,"Cape Verde",RB,,16.60707390,-24.20338430
3002,"Ribeira Grande",40,CV,"Cape Verde",RG,,37.82103690,-25.51481370
2984,"Ribeira Grande de Santiago",40,CV,"Cape Verde",RS,,14.98302980,-23.65617250
2998,Sal,40,CV,"Cape Verde",SL,,26.59581220,-80.20450830
2985,"Santa Catarina",40,CV,"Cape Verde",CA,,-27.24233920,-50.21885560
2995,"Santa Catarina do Fogo",40,CV,"Cape Verde",CF,,14.93091040,-24.32225770
3004,"Santa Cruz",40,CV,"Cape Verde",CR,,36.97411710,-122.03079630
2986,"São Domingos",40,CV,"Cape Verde",SD,,15.02861650,-23.56392200
3000,"São Filipe",40,CV,"Cape Verde",SF,,14.89516790,-24.49456360
2993,"São Lourenço dos Órgãos",40,CV,"Cape Verde",SO,,15.05378410,-23.60856120
2990,"São Miguel",40,CV,"Cape Verde",SM,,37.78041100,-25.49704660
3001,"São Vicente",40,CV,"Cape Verde",SV,,-23.96071570,-46.39620220
2992,"Sotavento Islands",40,CV,"Cape Verde",S,,15.00000000,-24.00000000
2983,Tarrafal,40,CV,"Cape Verde",TA,,15.27605780,-23.74840770
3003,"Tarrafal de São Nicolau",40,CV,"Cape Verde",TS,,16.56364980,-24.35494200
1259,"Bamingui-Bangoran Prefecture",42,CF,"Central African Republic",BB,,8.27334550,20.71224650
1262,Bangui,42,CF,"Central African Republic",BGF,,4.39467350,18.55818990
1264,"Basse-Kotto Prefecture",42,CF,"Central African Republic",BK,,4.87193190,21.28450250
1258,"Haut-Mbomou Prefecture",42,CF,"Central African Republic",HM,,6.25371340,25.47335540
1268,"Haute-Kotto Prefecture",42,CF,"Central African Republic",HK,,7.79643790,23.38235450
1263,"Kémo Prefecture",42,CF,"Central African Republic",KG,,5.88677940,19.37832060
1256,"Lobaye Prefecture",42,CF,"Central African Republic",LB,,4.35259810,17.47951730
1257,Mambéré-Kadéï,42,CF,"Central African Republic",HS,,4.70556530,15.96998780
1266,"Mbomou Prefecture",42,CF,"Central African Republic",MB,,5.55683700,23.76328280
1253,"Nana-Grébizi Economic Prefecture",42,CF,"Central African Republic",KB,,7.18486070,19.37832060
1260,"Nana-Mambéré Prefecture",42,CF,"Central African Republic",NM,,5.69321350,15.21948080
1255,"Ombella-M'Poko Prefecture",42,CF,"Central African Republic",MP,,5.11888250,18.42760470
1265,"Ouaka Prefecture",42,CF,"Central African Republic",UK,,6.31682160,20.71224650
1254,"Ouham Prefecture",42,CF,"Central African Republic",AC,,7.09091100,17.66888700
1267,"Ouham-Pendé Prefecture",42,CF,"Central African Republic",OP,,6.48509840,16.15809370
1252,Sangha-Mbaéré,42,CF,"Central African Republic",SE,,3.43686070,16.34637910
1261,"Vakaga Prefecture",42,CF,"Central African Republic",VK,,9.51132960,22.23840170
3583,"Bahr el Gazel",43,TD,Chad,BG,province,14.77022660,16.91225100
3590,Batha,43,TD,Chad,BA,province,13.93717750,18.42760470
3574,Borkou,43,TD,Chad,BO,province,17.86888450,18.80761950
5114,Chari-Baguirmi,43,TD,Chad,CB,province,11.46186260,15.24463940
3575,Ennedi-Est,43,TD,Chad,EE,province,16.34204960,23.00119890
3584,Ennedi-Ouest,43,TD,Chad,EO,province,18.97756300,21.85685860
3576,Guéra,43,TD,Chad,GR,province,11.12190150,18.42760470
3573,Hadjer-Lamis,43,TD,Chad,HL,province,12.45772730,16.72346390
3588,Kanem,43,TD,Chad,KA,province,14.87812620,15.40680790
3577,Lac,43,TD,Chad,LC,province,13.69153770,14.10013260
3585,"Logone Occidental",43,TD,Chad,LO,province,8.75967600,15.87600400
3591,"Logone Oriental",43,TD,Chad,LR,province,8.31499490,16.34637910
3589,Mandoul,43,TD,Chad,MA,province,8.60309100,17.47951730
3580,"Mayo-Kebbi Est",43,TD,Chad,ME,province,9.40460390,14.84546190
3571,"Mayo-Kebbi Ouest",43,TD,Chad,MO,province,10.41130140,15.59433880
3570,Moyen-Chari,43,TD,Chad,MC,province,9.06399980,18.42760470
3586,N'Djamena,43,TD,Chad,ND,province,12.13484570,15.05574150
3582,Ouaddaï,43,TD,Chad,OD,province,13.74847600,20.71224650
3592,Salamat,43,TD,Chad,SA,province,10.96916010,20.71224650
3572,Sila,43,TD,Chad,SI,province,12.13074000,21.28450250
3579,Tandjilé,43,TD,Chad,TA,province,9.66257290,16.72346390
3587,Tibesti,43,TD,Chad,TI,province,21.36500310,16.91225100
3581,"Wadi Fira",43,TD,Chad,WF,province,15.08924160,21.47528510
2828,"Aisén del General Carlos Ibañez del Campo",44,CL,Chile,AI,,-46.37834500,-72.30076230
2832,Antofagasta,44,CL,Chile,AN,,-23.83691040,-69.28775350
2829,"Arica y Parinacota",44,CL,Chile,AP,,-18.59404850,-69.47845410
2823,Atacama,44,CL,Chile,AT,,-27.56605580,-70.05031400
2827,Biobío,44,CL,Chile,BI,,-37.44644280,-72.14161320
2825,Coquimbo,44,CL,Chile,CO,,-30.54018100,-70.81199530
2826,"La Araucanía",44,CL,Chile,AR,,-38.94892100,-72.33111300
2838,"Libertador General Bernardo O'Higgins",44,CL,Chile,LI,,-34.57553740,-71.00223110
2835,"Los Lagos",44,CL,Chile,LL,,-41.91977790,-72.14161320
2834,"Los Ríos",44,CL,Chile,LR,,-40.23102170,-72.33111300
2836,"Magallanes y de la Antártica Chilena",44,CL,Chile,MA,,-52.20643160,-72.16850010
2833,Maule,44,CL,Chile,ML,,-35.51636030,-71.57239530
2831,Ñuble,44,CL,Chile,NB,,-36.72257430,-71.76224810
2824,"Región Metropolitana de Santiago",44,CL,Chile,RM,,-33.43755450,-70.65048960
2837,Tarapacá,44,CL,Chile,TA,,-20.20287990,-69.28775350
2830,Valparaíso,44,CL,Chile,VS,,-33.04723800,-71.61268850
2251,Anhui,45,CN,China,AH,province,30.60067730,117.92490020
2257,Beijing,45,CN,China,BJ,municipality,39.90419990,116.40739630
2271,Chongqing,45,CN,China,CQ,municipality,29.43158610,106.91225100
2248,Fujian,45,CN,China,FJ,province,26.48368420,117.92490020
2275,Gansu,45,CN,China,GS,province,35.75183260,104.28611160
2279,Guangdong,45,CN,China,GD,province,23.37903330,113.76328280
2278,"Guangxi Zhuang",45,CN,China,GX,"autonomous region",23.72475990,108.80761950
2261,Guizhou,45,CN,China,GZ,province,26.84296450,107.29028390
2273,Hainan,45,CN,China,HI,province,19.56639470,109.94968600
2280,Hebei,45,CN,China,HE,province,37.89565940,114.90422080
2265,Heilongjiang,45,CN,China,HL,province,47.12164720,128.73823100
2259,Henan,45,CN,China,HA,province,34.29043020,113.38235450
2267,"Hong Kong SAR",45,CN,China,HK,"special administrative region",22.31930390,114.16936110
2274,Hubei,45,CN,China,HB,province,30.73781180,112.23840170
2258,Hunan,45,CN,China,HN,province,27.36830090,109.28193470
2269,"Inner Mongolia",45,CN,China,NM,"autonomous region",43.37822000,115.05948150
2250,Jiangsu,45,CN,China,JS,province,33.14017150,119.78892480
2256,Jiangxi,45,CN,China,JX,province,27.08745640,114.90422080
2253,Jilin,45,CN,China,JL,province,43.83788300,126.54957200
2268,Liaoning,45,CN,China,LN,province,41.94365430,122.52903760
2266,"Macau SAR",45,CN,China,MO,"special administrative region",22.19874500,113.54387300
2262,"Ningxia Huizu",45,CN,China,NX,"autonomous region",37.19873100,106.15809370
2270,Qinghai,45,CN,China,QH,province,35.74479800,96.40773580
2272,Shaanxi,45,CN,China,SN,province,35.39399080,109.18800470
2252,Shandong,45,CN,China,SD,province,37.80060640,-122.26999180
2249,Shanghai,45,CN,China,SH,municipality,31.23041600,121.47370100
2254,Shanxi,45,CN,China,SX,province,37.24256490,111.85685860
2277,Sichuan,45,CN,China,SC,province,30.26380320,102.80547530
2255,Taiwan,45,CN,China,TW,province,23.69781000,120.96051500
2276,Tianjin,45,CN,China,TJ,municipality,39.12522910,117.01534350
2263,Xinjiang,45,CN,China,XJ,"autonomous region",42.52463570,87.53958550
2264,Xizang,45,CN,China,XZ,"autonomous region",30.15336050,88.78786780
2260,Yunnan,45,CN,China,YN,province,24.47528470,101.34310580
2247,Zhejiang,45,CN,China,ZJ,province,29.14164320,119.78892480
2895,Amazonas,48,CO,Colombia,AMA,,-1.44291230,-71.57239530
2890,Antioquia,48,CO,Colombia,ANT,,7.19860640,-75.34121790
2881,Arauca,48,CO,Colombia,ARA,,6.54730600,-71.00223110
2900,"Archipiélago de San Andrés, Providencia y Santa Catalina",48,CO,Colombia,SAP,,12.55673240,-81.71852530
2880,Atlántico,48,CO,Colombia,ATL,,10.69661590,-74.87410450
4921,"Bogotá D.C.",48,CO,Colombia,DC,"capital district",4.28204150,-74.50270420
2893,Bolívar,48,CO,Colombia,BOL,,8.67043820,-74.03001220
2903,Boyacá,48,CO,Colombia,BOY,,5.45451100,-73.36200300
2887,Caldas,48,CO,Colombia,CAL,,5.29826000,-75.24790610
2891,Caquetá,48,CO,Colombia,CAQ,,0.86989200,-73.84190630
2892,Casanare,48,CO,Colombia,CAS,,5.75892690,-71.57239530
2884,Cauca,48,CO,Colombia,CAU,,2.70498130,-76.82596520
2899,Cesar,48,CO,Colombia,CES,,9.33729480,-73.65362090
2876,Chocó,48,CO,Colombia,CHO,,5.25280330,-76.82596520
2898,Córdoba,48,CO,Colombia,COR,,8.04929300,-75.57405000
2875,Cundinamarca,48,CO,Colombia,CUN,,5.02600300,-74.03001220
2882,Guainía,48,CO,Colombia,GUA,,2.58539300,-68.52471490
2888,Guaviare,48,CO,Colombia,GUV,,2.04392400,-72.33111300
4871,Huila,48,CO,Colombia,HUI,department,2.53593490,-75.52766990
2889,"La Guajira",48,CO,Colombia,LAG,,11.35477430,-72.52048270
2886,Magdalena,48,CO,Colombia,MAG,,10.41130140,-74.40566120
2878,Meta,48,CO,Colombia,MET,,39.76732580,-104.97535950
2897,Nariño,48,CO,Colombia,NAR,,1.28915100,-77.35794000
2877,"Norte de Santander",48,CO,Colombia,NSA,,7.94628310,-72.89880690
2896,Putumayo,48,CO,Colombia,PUT,,0.43595060,-75.52766990
2874,Quindío,48,CO,Colombia,QUI,,4.46101910,-75.66735600
2879,Risaralda,48,CO,Colombia,RIS,,5.31584750,-75.99276520
2901,Santander,48,CO,Colombia,SAN,,6.64370760,-73.65362090
2902,Sucre,48,CO,Colombia,SUC,,8.81397700,-74.72328300
2883,Tolima,48,CO,Colombia,TOL,,4.09251680,-75.15453810
2904,"Valle del Cauca",48,CO,Colombia,VAC,,3.80088930,-76.64127120
2885,Vaupés,48,CO,Colombia,VAU,,0.85535610,-70.81199530
2894,Vichada,48,CO,Colombia,VID,,4.42344520,-69.28775350
2821,Anjouan,49,KM,Comoros,A,,-12.21381450,44.43706060
2822,"Grande Comore",49,KM,Comoros,G,,-11.71673380,43.36807880
2820,Mohéli,49,KM,Comoros,M,,-12.33773760,43.73340890
2866,"Bouenza Department",50,CG,Congo,11,,-4.11280790,13.72891670
2870,Brazzaville,50,CG,Congo,BZV,,-4.26335970,15.24288530
2864,"Cuvette Department",50,CG,Congo,8,,-0.28774460,16.15809370
2869,"Cuvette-Ouest Department",50,CG,Congo,15,,0.14475500,14.47233010
2867,"Kouilou Department",50,CG,Congo,5,,-4.14284130,11.88917210
2868,"Lékoumou Department",50,CG,Congo,2,,-3.17038200,13.35872880
2865,"Likouala Department",50,CG,Congo,7,,2.04392400,17.66888700
2872,"Niari Department",50,CG,Congo,9,,-3.18427000,12.25479190
2862,"Plateaux Department",50,CG,Congo,14,,-2.06800880,15.40680790
2863,Pointe-Noire,50,CG,Congo,16,,-4.76916230,11.86636200
2873,"Pool Department",50,CG,Congo,12,,-3.77626280,14.84546190
2871,"Sangha Department",50,CG,Congo,13,,1.46623280,15.40680790
1215,"Alajuela Province",53,CR,"Costa Rica",A,,10.39158300,-84.43827210
1209,"Guanacaste Province",53,CR,"Costa Rica",G,,10.62673990,-85.44367060
1212,"Heredia Province",53,CR,"Costa Rica",H,,10.47352300,-84.01674230
1213,"Limón Province",53,CR,"Costa Rica",L,,9.98963980,-83.03324170
1211,"Provincia de Cartago",53,CR,"Costa Rica",C,,9.86223110,-83.92141870
1210,"Puntarenas Province",53,CR,"Costa Rica",P,,9.21695310,-83.33618800
1214,"San José Province",53,CR,"Costa Rica",SJ,,9.91297270,-84.07682940
2634,Abidjan,54,CI,"Cote D'Ivoire (Ivory Coast)",AB,,5.35995170,-4.00825630
2626,Agnéby,54,CI,"Cote D'Ivoire (Ivory Coast)",16,,5.32245030,-4.34495290
2636,"Bafing Region",54,CI,"Cote D'Ivoire (Ivory Coast)",17,,8.32520470,-7.52472430
2643,"Bas-Sassandra District",54,CI,"Cote D'Ivoire (Ivory Coast)",BS,,5.27983560,-6.15269850
2635,"Bas-Sassandra Region",54,CI,"Cote D'Ivoire (Ivory Coast)",09,,5.35679160,-6.74939930
2654,"Comoé District",54,CI,"Cote D'Ivoire (Ivory Coast)",CM,,5.55279300,-3.25836260
2644,"Denguélé District",54,CI,"Cote D'Ivoire (Ivory Coast)",DN,,48.07077630,-68.56093410
2642,"Denguélé Region",54,CI,"Cote D'Ivoire (Ivory Coast)",10,,9.46623720,-7.43813550
2645,"Dix-Huit Montagnes",54,CI,"Cote D'Ivoire (Ivory Coast)",06,,7.37623730,-7.43813550
2633,Fromager,54,CI,"Cote D'Ivoire (Ivory Coast)",18,,45.54502130,-73.60462230
2651,"Gôh-Djiboua District",54,CI,"Cote D'Ivoire (Ivory Coast)",GD,,5.87113930,-5.56172790
2638,Haut-Sassandra,54,CI,"Cote D'Ivoire (Ivory Coast)",02,,6.87578480,-6.57833870
2632,"Lacs District",54,CI,"Cote D'Ivoire (Ivory Coast)",LC,,48.19801690,-80.45644120
2640,"Lacs Region",54,CI,"Cote D'Ivoire (Ivory Coast)",07,,47.73958660,-70.41866520
2627,"Lagunes District",54,CI,"Cote D'Ivoire (Ivory Coast)",LG,,5.88273340,-4.23333550
2639,"Lagunes region",54,CI,"Cote D'Ivoire (Ivory Coast)",01,,5.88273340,-4.23333550
2631,"Marahoué Region",54,CI,"Cote D'Ivoire (Ivory Coast)",12,,6.88462070,-5.89871390
2629,"Montagnes District",54,CI,"Cote D'Ivoire (Ivory Coast)",MG,,7.37623730,-7.43813550
2646,Moyen-Cavally,54,CI,"Cote D'Ivoire (Ivory Coast)",19,,6.52087930,-7.61142170
2630,Moyen-Comoé,54,CI,"Cote D'Ivoire (Ivory Coast)",05,,6.65149170,-3.50034540
2655,N'zi-Comoé,54,CI,"Cote D'Ivoire (Ivory Coast)",11,,7.24567490,-4.23333550
2648,"Sassandra-Marahoué District",54,CI,"Cote D'Ivoire (Ivory Coast)",SM,,6.88033480,-6.23759470
2625,"Savanes Region",54,CI,"Cote D'Ivoire (Ivory Coast)",03,,,
2628,Sud-Bandama,54,CI,"Cote D'Ivoire (Ivory Coast)",15,,5.53570830,-5.56172790
2652,Sud-Comoé,54,CI,"Cote D'Ivoire (Ivory Coast)",13,,5.55279300,-3.25836260
2637,"Vallée du Bandama District",54,CI,"Cote D'Ivoire (Ivory Coast)",VB,,8.27897800,-4.89356270
2647,"Vallée du Bandama Region",54,CI,"Cote D'Ivoire (Ivory Coast)",04,,8.27897800,-4.89356270
2650,"Woroba District",54,CI,"Cote D'Ivoire (Ivory Coast)",WR,,8.24913720,-6.92091350
2649,Worodougou,54,CI,"Cote D'Ivoire (Ivory Coast)",14,,8.25489620,-6.57833870
2653,Yamoussoukro,54,CI,"Cote D'Ivoire (Ivory Coast)",YM,,6.82762280,-5.28934330
2641,"Zanzan Region",54,CI,"Cote D'Ivoire (Ivory Coast)",ZZ,,8.82079040,-3.41955270
734,Bjelovar-Bilogora,55,HR,Croatia,07,county,45.89879720,16.84230930
737,Brod-Posavina,55,HR,Croatia,12,county,45.26379510,17.32645620
728,Dubrovnik-Neretva,55,HR,Croatia,19,county,43.07665880,17.52684710
743,Istria,55,HR,Croatia,18,county,45.12864550,13.90154200
5069,Karlovac,55,HR,Croatia,04,county,45.26133520,15.52542016
742,Koprivnica-Križevci,55,HR,Croatia,06,county,46.15689190,16.83908260
729,Krapina-Zagorje,55,HR,Croatia,02,county,46.10133930,15.88096930
731,Lika-Senj,55,HR,Croatia,09,county,44.61922180,15.47016080
726,Međimurje,55,HR,Croatia,20,county,46.37666440,16.42132980
740,Osijek-Baranja,55,HR,Croatia,14,county,45.55764280,18.39421410
724,Požega-Slavonia,55,HR,Croatia,11,county,45.34178680,17.81143590
735,"Primorje-Gorski Kotar",55,HR,Croatia,08,county,45.31739960,14.81674660
730,Šibenik-Knin,55,HR,Croatia,15,county,43.92814850,16.10376940
733,Sisak-Moslavina,55,HR,Croatia,03,county,45.38379260,16.53809940
725,Split-Dalmatia,55,HR,Croatia,17,county,43.52403280,16.81783770
739,Varaždin,55,HR,Croatia,05,county,46.23174730,16.33605590
732,Virovitica-Podravina,55,HR,Croatia,10,county,45.65579850,17.79324720
741,Vukovar-Syrmia,55,HR,Croatia,16,county,45.17735520,18.80535270
727,Zadar,55,HR,Croatia,13,county,44.14693900,15.61649430
736,Zagreb,55,HR,Croatia,01,county,45.87066120,16.39549100
738,Zagreb,55,HR,Croatia,21,city,45.81501080,15.98191890
283,"Artemisa Province",56,CU,Cuba,15,,22.75229030,-82.99316070
286,"Camagüey Province",56,CU,Cuba,09,,21.21672470,-77.74520810
282,"Ciego de Ávila Province",56,CU,Cuba,08,,21.93295150,-78.56608520
287,"Cienfuegos Province",56,CU,Cuba,06,,22.23797830,-80.36586500
275,"Granma Province",56,CU,Cuba,12,,20.38449020,-76.64127120
285,"Guantánamo Province",56,CU,Cuba,14,,20.14559170,-74.87410450
272,"Havana Province",56,CU,Cuba,03,,23.05406980,-82.34518900
279,"Holguín Province",56,CU,Cuba,11,,20.78378930,-75.80690820
278,"Isla de la Juventud",56,CU,Cuba,99,,21.70847370,-82.82202320
281,"Las Tunas Province",56,CU,Cuba,10,,21.06051620,-76.91820970
284,"Matanzas Province",56,CU,Cuba,04,,22.57671230,-81.33994140
276,"Mayabeque Province",56,CU,Cuba,16,,22.89265290,-81.95348150
277,"Pinar del Río Province",56,CU,Cuba,01,,22.40762560,-83.84730150
274,"Sancti Spíritus Province",56,CU,Cuba,07,,21.99382140,-79.47038850
273,"Santiago de Cuba Province",56,CU,Cuba,13,,20.23976820,-75.99276520
280,"Villa Clara Province",56,CU,Cuba,05,,22.49372040,-79.91927020
749,"Famagusta District (Mağusa)",57,CY,Cyprus,04,district,35.28570230,33.84112880
744,"Kyrenia District (Keryneia)",57,CY,Cyprus,06,district,35.29919400,33.23632460
747,"Larnaca District (Larnaka)",57,CY,Cyprus,03,district,34.85072060,33.48319060
748,"Limassol District (Leymasun)",57,CY,Cyprus,02,district,34.70713010,33.02261740
745,"Nicosia District (Lefkoşa)",57,CY,Cyprus,01,district,35.18556590,33.38227640
746,"Paphos District (Pafos)",57,CY,Cyprus,05,district,34.91645940,32.49200880
4627,Benešov,58,CZ,"Czech Republic",201,,49.69008280,14.77643990
4620,Beroun,58,CZ,"Czech Republic",202,,49.95734280,13.98407150
4615,Blansko,58,CZ,"Czech Republic",641,,49.36485020,16.64775520
4542,Břeclav,58,CZ,"Czech Republic",644,,48.75314000,16.88251690
4568,Brno-město,58,CZ,"Czech Republic",642,,49.19506020,16.60683710
4545,Brno-venkov,58,CZ,"Czech Republic",643,,49.12501380,16.45588240
4644,Bruntál,58,CZ,"Czech Republic",801,,49.98817670,17.46369410
4633,"Česká Lípa",58,CZ,"Czech Republic",511,,50.67852010,14.53969910
4556,"České Budějovice",58,CZ,"Czech Republic",311,,48.97755530,14.51507470
4543,"Český Krumlov",58,CZ,"Czech Republic",312,,48.81273540,14.31746570
4573,Cheb,58,CZ,"Czech Republic",411,,50.07953340,12.36986360
4553,Chomutov,58,CZ,"Czech Republic",422,,50.45838720,13.30179100
4634,Chrudim,58,CZ,"Czech Republic",531,,49.88302160,15.82908660
4609,Děčín,58,CZ,"Czech Republic",421,,50.77255630,14.21276120
4641,Domažlice,58,CZ,"Czech Republic",321,,49.43970270,12.93114350
4559,Frýdek-Místek,58,CZ,"Czech Republic",802,,49.68193050,18.36732160
4611,"Havlíčkův Brod",58,CZ,"Czech Republic",631,,49.60433640,15.57965520
4561,Hodonín,58,CZ,"Czech Republic",645,,48.85293910,17.12600250
4580,"Hradec Králové",58,CZ,"Czech Republic",521,,50.24148050,15.67430000
4612,"Jablonec nad Nisou",58,CZ,"Czech Republic",512,,50.72205280,15.17031350
4625,Jeseník,58,CZ,"Czech Republic",711,,50.22462490,17.19804710
4640,Jičín,58,CZ,"Czech Republic",522,,50.43533250,15.36104400
4613,Jihlava,58,CZ,"Czech Republic",632,,49.39837820,15.58704150
4639,"Jihočeský kraj",58,CZ,"Czech Republic",31,,48.94577890,14.44160550
4602,"Jihomoravský kraj",58,CZ,"Czech Republic",64,,48.95445280,16.76768990
4624,"Jindřichův Hradec",58,CZ,"Czech Republic",313,,49.14448230,15.00613890
4581,"Karlovarský kraj",58,CZ,"Czech Republic",41,,50.14350000,12.75018990
4604,"Karlovy Vary",58,CZ,"Czech Republic",412,,50.14350000,12.75018990
4586,Karviná,58,CZ,"Czech Republic",803,,49.85665240,18.54321860
4631,Kladno,58,CZ,"Czech Republic",203,,50.19402580,14.10436570
4591,Klatovy,58,CZ,"Czech Republic",322,,49.39555490,13.29509370
4618,Kolín,58,CZ,"Czech Republic",204,,49.98832930,15.05519770
4575,"Kraj Vysočina",58,CZ,"Czech Republic",63,,49.44900520,15.64059340
4614,"Královéhradecký kraj",58,CZ,"Czech Republic",52,,50.35124840,15.79764590
4593,Kroměříž,58,CZ,"Czech Republic",721,,49.29165820,17.39938000
4923,"Kutná Hora",58,CZ,"Czech Republic",205,,49.94920890,15.24704400
4590,Liberec,58,CZ,"Czech Republic",513,,50.75641010,14.99650410
4601,"Liberecký kraj",58,CZ,"Czech Republic",51,,50.65942400,14.76324240
4605,Litoměřice,58,CZ,"Czech Republic",423,,50.53841970,14.13054580
4617,Louny,58,CZ,"Czech Republic",424,,50.35398120,13.80335510
4638,Mělník,58,CZ,"Czech Republic",206,,50.31044150,14.51792230
4643,"Mladá Boleslav",58,CZ,"Czech Republic",207,,50.42523170,14.93624770
4600,"Moravskoslezský kraj",58,CZ,"Czech Republic",80,,49.73053270,18.23326370
4629,Most,58,CZ,"Czech Republic",425,,37.15540830,-94.29488840
4550,Náchod,58,CZ,"Czech Republic",523,,50.41457220,16.16563470
4548,"Nový Jičín",58,CZ,"Czech Republic",804,,49.59432510,18.01353560
4582,Nymburk,58,CZ,"Czech Republic",208,,50.18558160,15.04366040
4574,Olomouc,58,CZ,"Czech Republic",712,,49.59377800,17.25087870
4589,"Olomoucký kraj",58,CZ,"Czech Republic",71,,49.65865490,17.08114060
4623,Opava,58,CZ,"Czech Republic",805,,49.90837570,17.91633800
4584,Ostrava-město,58,CZ,"Czech Republic",806,,49.82092260,18.26252430
4547,Pardubice,58,CZ,"Czech Republic",532,,49.94444790,16.28569160
4588,"Pardubický kraj",58,CZ,"Czech Republic",53,,49.94444790,16.28569160
4645,Pelhřimov,58,CZ,"Czech Republic",633,,49.43062070,15.22298300
4560,Písek,58,CZ,"Czech Republic",314,,49.34199380,14.24697600
4608,Plzeň-jih,58,CZ,"Czech Republic",324,,49.59048850,13.57158610
4544,Plzeň-město,58,CZ,"Czech Republic",323,,49.73843140,13.37363710
4564,Plzeň-sever,58,CZ,"Czech Republic",325,,49.87748930,13.25374280
4607,"Plzeňský kraj",58,CZ,"Czech Republic",32,,49.41348120,13.31572460
4578,Prachatice,58,CZ,"Czech Republic",315,,49.01091000,14.00000050
4606,Praha-východ,58,CZ,"Czech Republic",209,,49.93893070,14.79244720
4619,Praha-západ,58,CZ,"Czech Republic",20A,,49.89352350,14.32937790
4598,"Praha, Hlavní město",58,CZ,"Czech Republic",10,,50.07553810,14.43780050
4626,Přerov,58,CZ,"Czech Republic",714,,49.46713560,17.50773320
4546,Příbram,58,CZ,"Czech Republic",20B,,49.69479590,14.08238100
4551,Prostějov,58,CZ,"Czech Republic",713,,49.44184010,17.12779040
4558,Rakovník,58,CZ,"Czech Republic",20C,,50.10612300,13.73966230
4583,Rokycany,58,CZ,"Czech Republic",326,,49.82628270,13.68749430
4636,"Rychnov nad Kněžnou",58,CZ,"Czech Republic",524,,50.16596510,16.27768420
4596,Semily,58,CZ,"Czech Republic",514,,50.60515760,15.32814090
4595,Sokolov,58,CZ,"Czech Republic",413,,50.20134340,12.60546360
4628,Strakonice,58,CZ,"Czech Republic",316,,49.26040430,13.91030850
4554,"Středočeský kraj",58,CZ,"Czech Republic",20,,49.87822230,14.93629550
4642,Šumperk,58,CZ,"Czech Republic",715,,49.97784070,16.97177540
4571,Svitavy,58,CZ,"Czech Republic",533,,49.75516290,16.46918610
4565,Tábor,58,CZ,"Czech Republic",317,,49.36462930,14.71912930
4646,Tachov,58,CZ,"Czech Republic",327,,49.79878030,12.63619210
4621,Teplice,58,CZ,"Czech Republic",426,,50.65846050,13.75132270
4597,Třebíč,58,CZ,"Czech Republic",634,,49.21478690,15.87955160
4579,Trutnov,58,CZ,"Czech Republic",525,,50.56538380,15.90909230
4592,"Uherské Hradiště",58,CZ,"Czech Republic",722,,49.05979690,17.49585010
4576,"Ústecký kraj",58,CZ,"Czech Republic",42,,50.61190370,13.78700860
4599,"Ústí nad Labem",58,CZ,"Czech Republic",427,,50.61190370,13.78700860
4647,"Ústí nad Orlicí",58,CZ,"Czech Republic",534,,49.97218010,16.39966170
4572,Vsetín,58,CZ,"Czech Republic",723,,49.37932500,18.06181620
4622,Vyškov,58,CZ,"Czech Republic",646,,49.21274450,16.98559270
4648,"Žďár nad Sázavou",58,CZ,"Czech Republic",635,,49.56430120,15.93910300
4563,Zlín,58,CZ,"Czech Republic",724,,49.16960520,17.80252200
4552,"Zlínský kraj",58,CZ,"Czech Republic",72,,49.21622960,17.77203530
4630,Znojmo,58,CZ,"Czech Republic",647,,48.92723270,16.10378080
2753,Bas-Uélé,51,CD,"Democratic Republic of the Congo",BU,,3.99010090,24.90422080
2744,Équateur,51,CD,"Democratic Republic of the Congo",EQ,,-1.83123900,-78.18340600
2750,Haut-Katanga,51,CD,"Democratic Republic of the Congo",HK,,-10.41020750,27.54958460
2758,Haut-Lomami,51,CD,"Democratic Republic of the Congo",HL,,-7.70527520,24.90422080
2734,Haut-Uélé,51,CD,"Democratic Republic of the Congo",HU,,3.58451540,28.29943500
2751,Ituri,51,CD,"Democratic Republic of the Congo",IT,,1.59576820,29.41793240
2757,Kasaï,51,CD,"Democratic Republic of the Congo",KS,,-5.04719790,20.71224650
2742,"Kasaï Central",51,CD,"Democratic Republic of the Congo",KC,,-8.44045910,20.41659340
2735,"Kasaï Oriental",51,CD,"Democratic Republic of the Congo",KE,,-6.03362300,23.57285010
2741,Kinshasa,51,CD,"Democratic Republic of the Congo",KN,,-4.44193110,15.26629310
2746,"Kongo Central",51,CD,"Democratic Republic of the Congo",BC,,-5.23656850,13.91439900
2740,Kwango,51,CD,"Democratic Republic of the Congo",KG,,-6.43374090,17.66888700
2759,Kwilu,51,CD,"Democratic Republic of the Congo",KL,,-5.11888250,18.42760470
2747,Lomami,51,CD,"Democratic Republic of the Congo",LO,,-6.14539310,24.52426400
4953,Lualaba,51,CD,"Democratic Republic of the Congo",LU,,-10.48086980,25.62978160
2755,Mai-Ndombe,51,CD,"Democratic Republic of the Congo",MN,,-2.63574340,18.42760470
2745,Maniema,51,CD,"Democratic Republic of the Congo",MA,,-3.07309290,26.04138890
2752,Mongala,51,CD,"Democratic Republic of the Congo",MO,,1.99623240,21.47528510
2749,Nord-Kivu,51,CD,"Democratic Republic of the Congo",NK,,-0.79177290,29.04599270
2739,Nord-Ubangi,51,CD,"Democratic Republic of the Congo",NU,,3.78787260,21.47528510
2743,Sankuru,51,CD,"Democratic Republic of the Congo",SA,,-2.84374530,23.38235450
2738,Sud-Kivu,51,CD,"Democratic Republic of the Congo",SK,,-3.01165800,28.29943500
2748,Sud-Ubangi,51,CD,"Democratic Republic of the Congo",SU,,3.22999420,19.18800470
2733,Tanganyika,51,CD,"Democratic Republic of the Congo",TA,,-6.27401180,27.92490020
2756,Tshopo,51,CD,"Democratic Republic of the Congo",TO,,0.54554620,24.90422080
2732,Tshuapa,51,CD,"Democratic Republic of the Congo",TU,,-0.99030230,23.02888440
1530,"Capital Region of Denmark",59,DK,Denmark,84,,55.67518120,12.54932610
1531,"Central Denmark Region",59,DK,Denmark,82,,56.30213900,9.30277700
1532,"North Denmark Region",59,DK,Denmark,81,,56.83074160,9.49305270
1529,"Region of Southern Denmark",59,DK,Denmark,83,,55.33077140,9.09249030
1528,"Region Zealand",59,DK,Denmark,85,,55.46325180,11.72149790
2933,"Ali Sabieh Region",60,DJ,Djibouti,AS,,11.19289730,42.94169800
2932,"Arta Region",60,DJ,Djibouti,AR,,11.52555280,42.84794740
2930,"Dikhil Region",60,DJ,Djibouti,DI,,11.10543360,42.37047440
2929,Djibouti,60,DJ,Djibouti,DJ,,11.82513800,42.59027500
2928,"Obock Region",60,DJ,Djibouti,OB,,12.38956910,43.01948970
2931,"Tadjourah Region",60,DJ,Djibouti,TA,,11.93388850,42.39383750
4082,"Saint Andrew Parish",61,DM,Dominica,02,,,
4078,"Saint David Parish",61,DM,Dominica,03,,,
4079,"Saint George Parish",61,DM,Dominica,04,,,
4076,"Saint John Parish",61,DM,Dominica,05,,,
4085,"Saint Joseph Parish",61,DM,Dominica,06,,39.02227120,-94.71765040
4083,"Saint Luke Parish",61,DM,Dominica,07,,42.10513630,-80.05707220
4077,"Saint Mark Parish",61,DM,Dominica,08,,,
4080,"Saint Patrick Parish",61,DM,Dominica,09,,,
4084,"Saint Paul Parish",61,DM,Dominica,10,,38.86146000,-90.74356190
4081,"Saint Peter Parish",61,DM,Dominica,11,,40.45241940,-80.00850560
4114,"Azua Province",62,DO,"Dominican Republic",02,,18.45527090,-70.73809280
4105,"Baoruco Province",62,DO,"Dominican Republic",03,,18.48798980,-71.41822490
4090,"Barahona Province",62,DO,"Dominican Republic",04,,18.21390660,-71.10437590
4107,"Dajabón Province",62,DO,"Dominican Republic",05,,19.54992410,-71.70865140
4095,"Distrito Nacional",62,DO,"Dominican Republic",01,,18.48605750,-69.93121170
4113,"Duarte Province",62,DO,"Dominican Republic",06,,19.20908230,-70.02700040
4086,"El Seibo Province",62,DO,"Dominican Republic",08,,18.76584960,-69.04066800
4102,"Espaillat Province",62,DO,"Dominican Republic",09,,19.62776580,-70.27867750
4106,"Hato Mayor Province",62,DO,"Dominican Republic",30,,18.76357990,-69.25576370
4089,"Hermanas Mirabal Province",62,DO,"Dominican Republic",19,,19.37475590,-70.35132350
4097,Independencia,62,DO,"Dominican Republic",10,,32.63357480,-115.42892940
4109,"La Altagracia Province",62,DO,"Dominican Republic",11,,18.58502360,-68.62010720
4087,"La Romana Province",62,DO,"Dominican Republic",12,,18.43102710,-68.98373730
4116,"La Vega Province",62,DO,"Dominican Republic",13,,19.22115540,-70.52887530
4094,"María Trinidad Sánchez Province",62,DO,"Dominican Republic",14,,19.37345970,-69.85144390
4099,"Monseñor Nouel Province",62,DO,"Dominican Republic",28,,18.92152340,-70.38368150
4115,"Monte Cristi Province",62,DO,"Dominican Republic",15,,19.73968990,-71.44339840
4111,"Monte Plata Province",62,DO,"Dominican Republic",29,,18.80808780,-69.78691460
4101,"Pedernales Province",62,DO,"Dominican Republic",16,,17.85376260,-71.33032090
4096,"Peravia Province",62,DO,"Dominican Republic",17,,18.27865940,-70.33358870
4092,"Puerto Plata Province",62,DO,"Dominican Republic",18,,19.75432250,-70.83328470
4103,"Samaná Province",62,DO,"Dominican Republic",20,,19.20583710,-69.33629490
4091,"San Cristóbal Province",62,DO,"Dominican Republic",21,,18.41804140,-70.10658490
4112,"San José de Ocoa Province",62,DO,"Dominican Republic",31,,18.54385800,-70.50418160
4098,"San Juan Province",62,DO,"Dominican Republic",22,,-31.52871270,-68.53604030
4110,"San Pedro de Macorís",62,DO,"Dominican Republic",23,,18.46266000,-69.30512340
4088,"Sánchez Ramírez Province",62,DO,"Dominican Republic",24,,19.05270600,-70.14922640
4108,"Santiago Province",62,DO,"Dominican Republic",25,,-33.45000000,-70.66670000
4100,"Santiago Rodríguez Province",62,DO,"Dominican Republic",26,,19.47131810,-71.33958010
4093,"Santo Domingo Province",62,DO,"Dominican Republic",32,,18.51042530,-69.84040540
4104,"Valverde Province",62,DO,"Dominican Republic",27,,19.58812210,-70.98033100
2923,Azuay,64,EC,Ecuador,A,province,-2.89430680,-78.99683440
2920,Bolívar,64,EC,Ecuador,B,province,-1.70958280,-79.04504290
2917,Cañar,64,EC,Ecuador,F,province,-2.55893150,-78.93881910
2915,Carchi,64,EC,Ecuador,C,province,0.50269120,-77.90425210
2925,Chimborazo,64,EC,Ecuador,H,province,-1.66479950,-78.65432550
2921,Cotopaxi,64,EC,Ecuador,X,province,-0.83842060,-78.66626780
2924,"El Oro",64,EC,Ecuador,O,province,-3.25924130,-79.95835410
2922,Esmeraldas,64,EC,Ecuador,E,province,0.96817890,-79.65172020
2905,Galápagos,64,EC,Ecuador,W,province,-0.95376910,-90.96560190
2914,Guayas,64,EC,Ecuador,G,province,-1.95748390,-79.91927020
2911,Imbabura,64,EC,Ecuador,I,province,0.34997680,-78.12601290
5068,Loja,64,EC,Ecuador,L,province,-3.99313000,-79.20422000
2910,"Los Ríos",64,EC,Ecuador,R,province,-1.02306070,-79.46088970
2913,Manabí,64,EC,Ecuador,M,province,-1.05434340,-80.45264400
2918,Morona-Santiago,64,EC,Ecuador,S,province,-2.30510620,-78.11468660
2916,Napo,64,EC,Ecuador,N,province,-0.99559640,-77.81296840
2926,Orellana,64,EC,Ecuador,D,province,-0.45451630,-76.99502860
2907,Pastaza,64,EC,Ecuador,Y,province,-1.48822650,-78.00310570
2927,Pichincha,64,EC,Ecuador,P,province,-0.14648470,-78.47519450
2912,"Santa Elena",64,EC,Ecuador,SE,province,-2.22671050,-80.85949900
2919,"Santo Domingo de los Tsáchilas",64,EC,Ecuador,SD,province,-0.25218820,-79.18793830
2906,Sucumbíos,64,EC,Ecuador,U,province,0.08892310,-76.88975570
2908,Tungurahua,64,EC,Ecuador,T,province,-1.26352840,-78.56608520
2909,"Zamora Chinchipe",64,EC,Ecuador,Z,province,-4.06558920,-78.95035250
3235,Alexandria,65,EG,Egypt,ALX,,30.87605680,29.74260400
3225,Aswan,65,EG,Egypt,ASN,,23.69664980,32.71813750
3236,Asyut,65,EG,Egypt,AST,,27.21338310,31.44561790
3241,Beheira,65,EG,Egypt,BH,,30.84809860,30.34355060
3230,"Beni Suef",65,EG,Egypt,BNS,,28.89388370,31.44561790
3223,Cairo,65,EG,Egypt,C,,29.95375640,31.53700030
3245,Dakahlia,65,EG,Egypt,DK,,31.16560440,31.49131820
3224,Damietta,65,EG,Egypt,DT,,31.36257990,31.67393710
3238,Faiyum,65,EG,Egypt,FYM,,29.30840210,30.84284970
3234,Gharbia,65,EG,Egypt,GH,,30.87535560,31.03351000
3239,Giza,65,EG,Egypt,GZ,,28.76662160,29.23207840
3244,Ismailia,65,EG,Egypt,IS,,30.58309340,32.26538870
3222,"Kafr el-Sheikh",65,EG,Egypt,KFS,,31.30854440,30.80394740
3242,Luxor,65,EG,Egypt,LX,,25.39444440,32.49200880
3231,Matrouh,65,EG,Egypt,MT,,29.56963500,26.41938900
3243,Minya,65,EG,Egypt,MN,,28.28472900,30.52790960
3228,Monufia,65,EG,Egypt,MNF,,30.59724550,30.98763210
3246,"New Valley",65,EG,Egypt,WAD,,24.54556380,27.17353160
3227,"North Sinai",65,EG,Egypt,SIN,,30.28236500,33.61757700
3229,"Port Said",65,EG,Egypt,PTS,,31.07586060,32.26538870
3232,Qalyubia,65,EG,Egypt,KB,,30.32923680,31.21684660
3247,Qena,65,EG,Egypt,KN,,26.23460330,32.98883190
3240,"Red Sea",65,EG,Egypt,BA,,24.68263160,34.15319470
5067,Sharqia,65,EG,Egypt,SHR,,30.67305450,31.15932470
3226,Sohag,65,EG,Egypt,SHG,,26.69383400,32.17460500
3237,"South Sinai",65,EG,Egypt,JS,,29.31018280,34.15319470
3233,Suez,65,EG,Egypt,SUZ,,29.36822550,32.17460500
4139,"Ahuachapán Department",66,SV,"El Salvador",AH,,13.82161480,-89.92532330
4132,"Cabañas Department",66,SV,"El Salvador",CA,,13.86482880,-88.74939980
4131,"Chalatenango Department",66,SV,"El Salvador",CH,,14.19166480,-89.17059980
4137,"Cuscatlán Department",66,SV,"El Salvador",CU,,13.86619570,-89.05615320
4134,"La Libertad Department",66,SV,"El Salvador",LI,,13.68176610,-89.36062980
4136,"La Paz Department",66,SV,"El Salvador",PA,,,
4138,"La Unión Department",66,SV,"El Salvador",UN,,13.48864430,-87.89424510
4130,"Morazán Department",66,SV,"El Salvador",MO,,13.76820000,-88.12913870
4135,"San Miguel Department",66,SV,"El Salvador",SM,,13.44510410,-88.24611830
4133,"San Salvador Department",66,SV,"El Salvador",SS,,13.77399970,-89.20867730
4127,"San Vicente Department",66,SV,"El Salvador",SV,,13.58685610,-88.74939980
4128,"Santa Ana Department",66,SV,"El Salvador",SA,,14.14611210,-89.51200840
4140,"Sonsonate Department",66,SV,"El Salvador",SO,,13.68235800,-89.66281110
4129,"Usulután Department",66,SV,"El Salvador",US,,13.44706340,-88.55653100
3444,"Annobón Province",67,GQ,"Equatorial Guinea",AN,,-1.42687820,5.63528010
3446,"Bioko Norte Province",67,GQ,"Equatorial Guinea",BN,,3.65950720,8.79218360
3443,"Bioko Sur Province",67,GQ,"Equatorial Guinea",BS,,3.42097850,8.61606740
3445,"Centro Sur Province",67,GQ,"Equatorial Guinea",CS,,1.34360840,10.43965600
3442,"Insular Region",67,GQ,"Equatorial Guinea",I,,37.09024000,-95.71289100
3439,"Kié-Ntem Province",67,GQ,"Equatorial Guinea",KN,,2.02809300,11.07117580
3441,"Litoral Province",67,GQ,"Equatorial Guinea",LI,,1.57502440,9.81249350
3438,"Río Muni",67,GQ,"Equatorial Guinea",C,,1.46106060,9.67868940
3440,"Wele-Nzas Province",67,GQ,"Equatorial Guinea",WN,,1.41661620,11.07117580
3425,"Anseba Region",68,ER,Eritrea,AN,,16.47455310,37.80876930
3427,"Debub Region",68,ER,Eritrea,DU,,14.94786920,39.15436770
3428,"Gash-Barka Region",68,ER,Eritrea,GB,,15.40688250,37.63866220
3426,"Maekel Region",68,ER,Eritrea,MA,,15.35514090,38.86236830
3424,"Northern Red Sea Region",68,ER,Eritrea,SK,,16.25839970,38.82054540
3429,"Southern Red Sea Region",68,ER,Eritrea,DK,,13.51371030,41.76064720
3567,"Harju County",69,EE,Estonia,37,,59.33342390,25.24669740
3555,"Hiiu County",69,EE,Estonia,39,,58.92395530,22.59194680
3569,"Ida-Viru County",69,EE,Estonia,44,,59.25926630,27.41365350
3566,"Järva County",69,EE,Estonia,51,,58.88667130,25.50006240
3565,"Jõgeva County",69,EE,Estonia,49,,58.75061430,26.36048780
3568,"Lääne County",69,EE,Estonia,57,,58.97227420,23.87408340
3564,"Lääne-Viru County",69,EE,Estonia,59,,59.30188160,26.32803120
3562,"Pärnu County",69,EE,Estonia,67,,58.52619520,24.40201590
3563,"Põlva County",69,EE,Estonia,65,,58.11606220,27.20663940
3559,"Rapla County",69,EE,Estonia,70,,58.84926250,24.73465690
3561,"Saare County",69,EE,Estonia,74,,58.48497210,22.61364080
3557,"Tartu County",69,EE,Estonia,78,,58.40571280,26.80157600
3558,"Valga County",69,EE,Estonia,82,,57.91034410,26.16018190
3556,"Viljandi County",69,EE,Estonia,84,,58.28217460,25.57522330
3560,"Võru County",69,EE,Estonia,86,,57.73773720,27.13989380
969,"Hhohho District",212,SZ,Eswatini,HH,,-26.13656620,31.35416310
970,"Lubombo District",212,SZ,Eswatini,LU,,-26.78517730,31.81070790
968,"Manzini District",212,SZ,Eswatini,MA,,-26.50819990,31.37131640
971,"Shiselweni District",212,SZ,Eswatini,SH,,-26.98275770,31.35416310
11,"Addis Ababa",70,ET,Ethiopia,AA,,8.98060340,38.75776050
6,"Afar Region",70,ET,Ethiopia,AF,,11.75593880,40.95868800
3,"Amhara Region",70,ET,Ethiopia,AM,,11.34942470,37.97845850
9,"Benishangul-Gumuz Region",70,ET,Ethiopia,BE,,10.78028890,35.56578620
8,"Dire Dawa",70,ET,Ethiopia,DD,,9.60087470,41.85014200
10,"Gambela Region",70,ET,Ethiopia,GA,,7.92196870,34.15319470
7,"Harari Region",70,ET,Ethiopia,HA,,9.31486600,42.19677160
5,"Oromia Region",70,ET,Ethiopia,OR,,7.54603770,40.63468510
2,"Somali Region",70,ET,Ethiopia,SO,,6.66122930,43.79084530
1,"Southern Nations, Nationalities, and Peoples' Region",70,ET,Ethiopia,SN,,6.51569110,36.95410700
4,"Tigray Region",70,ET,Ethiopia,TI,,14.03233360,38.31657250
1917,Ba,73,FJ,"Fiji Islands",01,,36.06138930,-95.80058720
1930,Bua,73,FJ,"Fiji Islands",02,,43.09645840,-89.50088000
1924,Cakaudrove,73,FJ,"Fiji Islands",03,,-16.58141050,179.51200840
1929,"Central Division",73,FJ,"Fiji Islands",C,,34.04400660,-118.24727380
1932,"Eastern Division",73,FJ,"Fiji Islands",E,,32.80943050,-117.12899370
1934,Kadavu,73,FJ,"Fiji Islands",04,,-19.01271220,178.18766760
1933,Lau,73,FJ,"Fiji Islands",05,,31.66870150,-106.39557630
1916,Lomaiviti,73,FJ,"Fiji Islands",06,,-17.70900000,179.09100000
1922,Macuata,73,FJ,"Fiji Islands",07,,-16.48649220,179.28472510
1919,Nadroga-Navosa,73,FJ,"Fiji Islands",08,,-17.98652780,177.65811300
1927,Naitasiri,73,FJ,"Fiji Islands",09,,-17.89757540,178.20715980
1928,Namosi,73,FJ,"Fiji Islands",10,,-18.08641760,178.12913870
1921,"Northern Division",73,FJ,"Fiji Islands",N,,32.87687660,-117.21563450
1926,Ra,73,FJ,"Fiji Islands",11,,37.10031530,-95.67442460
1920,Rewa,73,FJ,"Fiji Islands",12,,34.79235170,-82.36092640
1931,Rotuma,73,FJ,"Fiji Islands",R,,-12.50250690,177.07241640
1925,Serua,73,FJ,"Fiji Islands",13,,-18.18047490,178.05097900
1918,Tailevu,73,FJ,"Fiji Islands",14,,-17.82691110,178.29324800
1923,"Western Division",73,FJ,"Fiji Islands",W,,42.96621980,-78.70211340
1509,"Åland Islands",74,FI,Finland,01,region,60.17852470,19.91561050
1511,"Central Finland",74,FI,Finland,08,region,62.56667430,25.55494450
1494,"Central Ostrobothnia",74,FI,Finland,07,region,63.56217350,24.00136310
1507,"Finland Proper",74,FI,Finland,19,region,60.36279140,22.44393690
1496,Kainuu,74,FI,Finland,05,region,64.37365640,28.74374750
1512,Kymenlaakso,74,FI,Finland,09,region,60.78051200,26.88293360
1500,Lapland,74,FI,Finland,10,region,67.92223040,26.50464380
1504,"North Karelia",74,FI,Finland,13,region,62.80620780,30.15538870
1505,"Northern Ostrobothnia",74,FI,Finland,14,region,65.27949300,26.28904170
1503,"Northern Savonia",74,FI,Finland,15,region,63.08448000,27.02535040
1508,Ostrobothnia,74,FI,Finland,12,region,63.11817570,21.90610620
1502,"Päijänne Tavastia",74,FI,Finland,16,region,61.32300410,25.73224960
1506,Pirkanmaa,74,FI,Finland,11,region,61.69869180,23.78955980
1501,Satakunta,74,FI,Finland,17,region,61.59327580,22.14830810
1497,"South Karelia",74,FI,Finland,02,region,61.11819490,28.10243720
1498,"Southern Ostrobothnia",74,FI,Finland,03,region,62.94330990,23.52852670
1495,"Southern Savonia",74,FI,Finland,04,region,61.69451480,27.80050150
1493,"Tavastia Proper",74,FI,Finland,06,region,60.90701500,24.30054980
1510,Uusimaa,74,FI,Finland,18,region,60.21872000,25.27162100
4967,Ain,75,FR,France,01,"metropolitan department",46.06508600,4.88861500
4968,Aisne,75,FR,France,02,"metropolitan department",49.45289210,3.04651110
4969,Allier,75,FR,France,03,"metropolitan department",46.36708630,2.58082770
4970,Alpes-de-Haute-Provence,75,FR,France,04,"metropolitan department",44.16377520,5.67247800
4972,Alpes-Maritimes,75,FR,France,06,"metropolitan department",43.92041700,6.61678220
4811,Alsace,75,FR,France,6AE,"European collectivity",48.31817950,7.44162410
4973,Ardèche,75,FR,France,07,"metropolitan department",44.81486950,3.81334830
4974,Ardennes,75,FR,France,08,"metropolitan department",49.69759510,4.14895760
4975,Ariège,75,FR,France,09,"metropolitan department",42.94347830,0.94048640
4976,Aube,75,FR,France,10,"metropolitan department",48.31975470,3.56371040
4977,Aude,75,FR,France,11,"metropolitan department",43.05411400,1.90384760
4798,Auvergne-Rhône-Alpes,75,FR,France,ARA,"metropolitan region",45.44714310,4.38525070
4978,Aveyron,75,FR,France,12,"metropolitan department",44.31563620,2.08523790
5035,Bas-Rhin,75,FR,France,67,"metropolitan department",48.59864440,7.02666760
4979,Bouches-du-Rhône,75,FR,France,13,"metropolitan department",43.54038650,4.46138290
4825,Bourgogne-Franche-Comté,75,FR,France,BFC,"metropolitan region",47.28051270,4.99943720
4807,Bretagne,75,FR,France,BRE,"metropolitan region",48.20204710,-2.93264350
4981,Calvados,75,FR,France,14,"metropolitan department",49.09035140,-0.91706480
4982,Cantal,75,FR,France,15,"metropolitan department",45.04921770,2.15672720
4818,"Centre-Val de Loire",75,FR,France,CVL,"metropolitan region",47.75156860,1.67506310
4983,Charente,75,FR,France,16,"metropolitan department",45.66584790,-0.31845770
4984,Charente-Maritime,75,FR,France,17,"metropolitan department",45.72968280,-1.33881160
4985,Cher,75,FR,France,18,"metropolitan department",47.02436280,1.86627320
5064,Clipperton,75,FR,France,CP,dependency,10.28335410,-109.22542150
4986,Corrèze,75,FR,France,19,"metropolitan department",45.34237070,1.31717330
4806,Corse,75,FR,France,20R,"metropolitan collectivity with special status",42.03960420,9.01289260
4996,Corse-du-Sud,75,FR,France,2A,"metropolitan department",41.85720550,8.41091830
4987,Côte-d'Or,75,FR,France,21,"metropolitan department",47.46513020,4.23154950
4988,Côtes-d'Armor,75,FR,France,22,"metropolitan department",48.46633360,-3.34789610
4989,Creuse,75,FR,France,23,"metropolitan department",46.05903940,1.43150500
5047,Deux-Sèvres,75,FR,France,79,"metropolitan department",46.53868170,-0.90199480
4990,Dordogne,75,FR,France,24,"metropolitan department",45.14234160,0.14274080
4991,Doubs,75,FR,France,25,"metropolitan department",46.93217740,6.34762140
4992,Drôme,75,FR,France,26,"metropolitan department",44.72933570,4.67821580
5059,Essonne,75,FR,France,91,"metropolitan department",48.53046150,1.96990560
4993,Eure,75,FR,France,27,"metropolitan department",49.07540350,0.48937320
4994,Eure-et-Loir,75,FR,France,28,"metropolitan department",48.44697840,0.81470250
4995,Finistère,75,FR,France,29,"metropolitan department",48.22696100,-4.82437330
4822,"French Guiana",75,FR,France,973,"overseas region",3.93388900,-53.12578200
4824,"French Polynesia",75,FR,France,PF,"overseas collectivity",-17.67974200,-149.40684300
5065,"French Southern and Antarctic Lands",75,FR,France,TF,"overseas territory",-47.54466040,51.28375420
4998,Gard,75,FR,France,30,"metropolitan department",43.95952760,3.49356810
5000,Gers,75,FR,France,32,"metropolitan department",43.69505340,-0.09997280
5001,Gironde,75,FR,France,33,"metropolitan department",44.89584690,-1.59405320
4820,Grand-Est,75,FR,France,GES,"metropolitan region",48.69980300,6.18780740
4829,Guadeloupe,75,FR,France,971,"overseas region",16.26500000,-61.55100000
5036,Haut-Rhin,75,FR,France,68,"metropolitan department",47.86537740,6.67113810
4997,Haute-Corse,75,FR,France,2B,"metropolitan department",42.42958660,8.50625610
4999,Haute-Garonne,75,FR,France,31,"metropolitan department",43.30505550,0.68455150
5011,Haute-Loire,75,FR,France,43,"metropolitan department",45.08538060,3.22607070
5020,Haute-Marne,75,FR,France,52,"metropolitan department",48.13248210,4.69834990
5039,Haute-Saône,75,FR,France,70,"metropolitan department",47.63789960,5.53550550
5043,Haute-Savoie,75,FR,France,74,"metropolitan department",46.04452770,5.86413800
5055,Haute-Vienne,75,FR,France,87,"metropolitan department",45.91868780,0.70972060
4971,Hautes-Alpes,75,FR,France,05,"metropolitan department",44.65626820,5.68732110
5033,Hautes-Pyrénées,75,FR,France,65,"metropolitan department",43.14294620,-0.40097360
4828,Hauts-de-France,75,FR,France,HDF,"metropolitan region",50.48011530,2.79372650
5060,Hauts-de-Seine,75,FR,France,92,"metropolitan department",48.84030080,2.10125590
5002,Hérault,75,FR,France,34,"metropolitan department",43.59111200,2.80661080
4796,Île-de-France,75,FR,France,IDF,"metropolitan region",48.84991980,2.63704110
5003,Ille-et-Vilaine,75,FR,France,35,"metropolitan department",48.17624840,-2.21304010
5004,Indre,75,FR,France,36,"metropolitan department",46.81175500,0.97555230
5005,Indre-et-Loire,75,FR,France,37,"metropolitan department",47.22285820,0.14896190
5006,Isère,75,FR,France,38,"metropolitan department",45.28922710,4.99023550
5007,Jura,75,FR,France,39,"metropolitan department",46.78287410,5.16918440
4823,"La Réunion",75,FR,France,974,"overseas region",-21.11514100,55.53638400
5008,Landes,75,FR,France,40,"metropolitan department",44.00950800,-1.25385790
5009,Loir-et-Cher,75,FR,France,41,"metropolitan department",47.65937600,0.85376310
5010,Loire,75,FR,France,42,"metropolitan department",46.35228120,-1.17563390
5012,Loire-Atlantique,75,FR,France,44,"metropolitan department",47.34757210,-2.34663120
5013,Loiret,75,FR,France,45,"metropolitan department",47.91354310,1.76009900
5014,Lot,75,FR,France,46,"metropolitan department",44.62460700,1.03576310
5015,Lot-et-Garonne,75,FR,France,47,"metropolitan department",44.36873140,-0.09161690
5016,Lozère,75,FR,France,48,"metropolitan department",44.54227790,2.92934590
5017,Maine-et-Loire,75,FR,France,49,"metropolitan department",47.38900340,-1.12025270
5018,Manche,75,FR,France,50,"metropolitan department",49.08817340,-2.46272090
5019,Marne,75,FR,France,51,"metropolitan department",48.96107450,3.65737670
4827,Martinique,75,FR,France,972,"overseas region",14.64152800,-61.02417400
5021,Mayenne,75,FR,France,53,"metropolitan department",48.30668420,-0.64901820
4797,Mayotte,75,FR,France,976,"overseas region",-12.82750000,45.16624400
5038,"Métropole de Lyon",75,FR,France,69M,"metropolitan department",45.74826290,4.59584040
5022,Meurthe-et-Moselle,75,FR,France,54,"metropolitan department",48.95566150,5.71423500
5023,Meuse,75,FR,France,55,"metropolitan department",49.01246200,4.81087340
5024,Morbihan,75,FR,France,56,"metropolitan department",47.74395180,-3.44555240
5025,Moselle,75,FR,France,57,"metropolitan department",49.02045660,6.20553220
5026,Nièvre,75,FR,France,58,"metropolitan department",47.11921640,2.97797130
5027,Nord,75,FR,France,59,"metropolitan department",50.52854770,2.60007760
4804,Normandie,75,FR,France,NOR,"metropolitan region",48.87987040,0.17125290
4795,Nouvelle-Aquitaine,75,FR,France,NAQ,"metropolitan region",45.70871820,0.62689100
4799,Occitanie,75,FR,France,OCC,"metropolitan region",43.89272320,3.28276250
5028,Oise,75,FR,France,60,"metropolitan department",49.41173350,1.86688250
5029,Orne,75,FR,France,61,"metropolitan department",48.57576440,-0.50242950
4816,Paris,75,FR,France,75C,"metropolitan collectivity with special status",48.85661400,2.35222190
5030,Pas-de-Calais,75,FR,France,62,"metropolitan department",50.51446990,1.81149800
4802,Pays-de-la-Loire,75,FR,France,PDL,"metropolitan region",47.76328360,-0.32996870
4812,Provence-Alpes-Côte-d’Azur,75,FR,France,PAC,"metropolitan region",43.93516910,6.06791940
5031,Puy-de-Dôme,75,FR,France,63,"metropolitan department",45.77141850,2.62626760
5032,Pyrénées-Atlantiques,75,FR,France,64,"metropolitan department",43.18681700,-1.44170710
5034,Pyrénées-Orientales,75,FR,France,66,"metropolitan department",42.62541790,1.88929580
5037,Rhône,75,FR,France,69,"metropolitan department",44.93433000,4.24093290
4821,"Saint Pierre and Miquelon",75,FR,France,PM,"overseas collectivity",46.88520000,-56.31590000
4794,Saint-Barthélemy,75,FR,France,BL,"overseas collectivity",17.90051340,-62.82058710
4809,Saint-Martin,75,FR,France,MF,"overseas collectivity",18.07082980,-63.05008090
5040,Saône-et-Loire,75,FR,France,71,"metropolitan department",46.65548830,3.98350500
5041,Sarthe,75,FR,France,72,"metropolitan department",48.02627330,-0.32613170
5042,Savoie,75,FR,France,73,"metropolitan department",45.49469900,5.84329840
5045,Seine-et-Marne,75,FR,France,77,"metropolitan department",48.61853940,2.41525610
5044,Seine-Maritime,75,FR,France,76,"metropolitan department",49.66096810,0.36775610
5061,Seine-Saint-Denis,75,FR,France,93,"metropolitan department",48.90993180,2.30573790
5048,Somme,75,FR,France,80,"metropolitan department",49.96859220,1.73106960
5049,Tarn,75,FR,France,81,"metropolitan department",43.79149770,1.67588930
5050,Tarn-et-Garonne,75,FR,France,82,"metropolitan department",44.08089500,1.08916570
5058,"Territoire de Belfort",75,FR,France,90,"metropolitan department",47.62930720,6.66962000
5063,Val-d'Oise,75,FR,France,95,"metropolitan department",49.07518180,1.82169140
5062,Val-de-Marne,75,FR,France,94,"metropolitan department",48.77470040,2.32210390
5051,Var,75,FR,France,83,"metropolitan department",43.39507300,5.73424170
5052,Vaucluse,75,FR,France,84,"metropolitan department",44.04475000,4.64277180
5053,Vendée,75,FR,France,85,"metropolitan department",46.67541030,-2.02983920
5054,Vienne,75,FR,France,86,"metropolitan department",45.52213140,4.84531360
5056,Vosges,75,FR,France,88,"metropolitan department",48.16301730,5.73556000
4810,"Wallis and Futuna",75,FR,France,WF,"overseas collectivity",-14.29380000,-178.11650000
5057,Yonne,75,FR,France,89,"metropolitan department",47.85476140,3.03394040
5046,Yvelines,75,FR,France,78,"metropolitan department",48.76153010,1.27729490
2727,"Estuaire Province",79,GA,Gabon,1,,0.44328640,10.08072980
2726,"Haut-Ogooué Province",79,GA,Gabon,2,,-1.47625440,13.91439900
2730,"Moyen-Ogooué Province",79,GA,Gabon,3,,-0.44278400,10.43965600
2731,"Ngounié Province",79,GA,Gabon,4,,-1.49303030,10.98070030
2725,"Nyanga Province",79,GA,Gabon,5,,-2.88210330,11.16173560
2724,"Ogooué-Ivindo Province",79,GA,Gabon,6,,0.88183110,13.17403480
2729,"Ogooué-Lolo Province",79,GA,Gabon,7,,-0.88440930,12.43805810
2728,"Ogooué-Maritime Province",79,GA,Gabon,8,,-1.34659750,9.72326730
2723,"Woleu-Ntem Province",79,GA,Gabon,9,,2.29898270,11.44669140
2666,Banjul,80,GM,"Gambia The",B,,13.45487610,-16.57903230
2669,"Central River Division",80,GM,"Gambia The",M,,13.59944690,-14.89216680
2670,"Lower River Division",80,GM,"Gambia The",L,,13.35533060,-15.92299000
2671,"North Bank Division",80,GM,"Gambia The",N,,13.52854360,-16.01699710
2668,"Upper River Division",80,GM,"Gambia The",U,,13.42573660,-14.00723480
2667,"West Coast Division",80,GM,"Gambia The",W,,5.97727980,116.07542880
901,Abkhazia,81,GE,Georgia,AB,,43.00155440,41.02340700
900,Adjara,81,GE,Georgia,AJ,,41.60056260,42.06883830
907,Guria,81,GE,Georgia,GU,,41.94427360,42.04580910
905,Imereti,81,GE,Georgia,IM,,42.23010800,42.90086640
910,Kakheti,81,GE,Georgia,KA,,41.64816020,45.69055540
897,"Khelvachauri Municipality",81,GE,Georgia,29,,41.58019260,41.66107420
904,"Kvemo Kartli",81,GE,Georgia,KK,,41.47918330,44.65604510
902,Mtskheta-Mtianeti,81,GE,Georgia,MM,,42.16821850,44.65060580
909,"Racha-Lechkhumi and Kvemo Svaneti",81,GE,Georgia,RL,,42.67188730,43.05628360
908,"Samegrelo-Zemo Svaneti",81,GE,Georgia,SZ,,42.73522470,42.16893620
906,Samtskhe-Javakheti,81,GE,Georgia,SJ,,41.54792960,43.27764000
898,"Senaki Municipality",81,GE,Georgia,50,,42.26963600,42.06568960
903,"Shida Kartli",81,GE,Georgia,SK,,42.07569440,43.95404620
899,Tbilisi,81,GE,Georgia,TB,,41.71513770,44.82709600
3006,Baden-Württemberg,82,DE,Germany,BW,,48.66160370,9.35013360
3009,Bavaria,82,DE,Germany,BY,,48.79044720,11.49788950
3010,Berlin,82,DE,Germany,BE,,52.52000660,13.40495400
3013,Brandenburg,82,DE,Germany,BB,,52.41252870,12.53164440
3014,Bremen,82,DE,Germany,HB,,53.07929620,8.80169360
3016,Hamburg,82,DE,Germany,HH,,53.55108460,9.99368190
3018,Hesse,82,DE,Germany,HE,,50.65205150,9.16243760
3008,"Lower Saxony",82,DE,Germany,NI,,52.63670360,9.84507660
3007,Mecklenburg-Vorpommern,82,DE,Germany,MV,,53.61265050,12.42959530
3017,"North Rhine-Westphalia",82,DE,Germany,NW,,51.43323670,7.66159380
3019,Rhineland-Palatinate,82,DE,Germany,RP,,50.11834600,7.30895270
3020,Saarland,82,DE,Germany,SL,,49.39642340,7.02296070
3021,Saxony,82,DE,Germany,SN,,51.10454070,13.20173840
3011,Saxony-Anhalt,82,DE,Germany,ST,,51.95026490,11.69227340
3005,Schleswig-Holstein,82,DE,Germany,SH,,54.21936720,9.69611670
3015,Thuringia,82,DE,Germany,TH,,51.01098920,10.84534600
53,Ahafo,83,GH,Ghana,AF,region,7.58213720,-2.54974630
48,Ashanti,83,GH,Ghana,AH,region,6.74704360,-1.52086240
4959,Bono,83,GH,Ghana,BO,region,7.65000000,-2.50000000
4958,"Bono East",83,GH,Ghana,BE,region,7.75000000,-1.05000000
52,Central,83,GH,Ghana,CP,region,5.50000000,-1.00000000
50,Eastern,83,GH,Ghana,EP,region,6.50000000,-0.50000000
54,"Greater Accra",83,GH,Ghana,AA,region,5.81428360,0.07467670
4960,"North East",83,GH,Ghana,NE,region,10.51666700,-0.36666700
51,Northern,83,GH,Ghana,NP,region,9.50000000,-1.00000000
4961,Oti,83,GH,Ghana,OT,region,7.90000000,0.30000000
4962,Savannah,83,GH,Ghana,SV,region,9.08333300,-1.81666700
55,"Upper East",83,GH,Ghana,UE,region,10.70824990,-0.98206680
57,"Upper West",83,GH,Ghana,UW,region,10.25297570,-2.14502450
56,Volta,83,GH,Ghana,TV,region,6.57813730,0.45023680
49,Western,83,GH,Ghana,WP,region,5.50000000,-2.50000000
4963,"Western North",83,GH,Ghana,WN,region,6.30000000,-2.80000000
2116,"Achaea Regional Unit",85,GR,Greece,13,,38.11587290,21.95224910
2123,"Aetolia-Acarnania Regional Unit",85,GR,Greece,01,,38.70843860,21.37989280
2098,"Arcadia Prefecture",85,GR,Greece,12,,37.55578250,22.33377690
2105,"Argolis Regional Unit",85,GR,Greece,11,,,
2122,"Attica Region",85,GR,Greece,I,,38.04575680,23.85847370
2126,"Boeotia Regional Unit",85,GR,Greece,03,,38.36636640,23.09650640
2128,"Central Greece Region",85,GR,Greece,H,,38.60439840,22.71521310
2125,"Central Macedonia",85,GR,Greece,B,,40.62117300,23.19180210
2115,"Chania Regional Unit",85,GR,Greece,94,,35.51382980,24.01803670
2124,"Corfu Prefecture",85,GR,Greece,22,,39.62498380,19.92234610
2129,"Corinthia Regional Unit",85,GR,Greece,15,,,
2109,"Crete Region",85,GR,Greece,M,,35.24011700,24.80926910
2130,"Drama Regional Unit",85,GR,Greece,52,,41.23400230,24.23904980
2120,"East Attica Regional Unit",85,GR,Greece,A2,,38.20540930,23.85847370
2117,"East Macedonia and Thrace",85,GR,Greece,A,,41.12951260,24.88771910
2110,"Epirus Region",85,GR,Greece,D,,39.57064130,20.76428430
2101,Euboea,85,GR,Greece,04,,38.52360360,23.85847370
2102,"Grevena Prefecture",85,GR,Greece,51,,40.08376260,21.42732990
2099,"Imathia Regional Unit",85,GR,Greece,53,,40.60600670,22.14302150
2113,"Ioannina Regional Unit",85,GR,Greece,33,,39.66502880,20.85374660
2131,"Ionian Islands Region",85,GR,Greece,F,,37.96948980,21.38023720
2095,"Karditsa Regional Unit",85,GR,Greece,41,,39.36402580,21.92140490
2100,"Kastoria Regional Unit",85,GR,Greece,56,,40.51926910,21.26871710
2127,"Kefalonia Prefecture",85,GR,Greece,23,,38.17536750,20.56921790
2111,"Kilkis Regional Unit",85,GR,Greece,57,,40.99370710,22.87536740
2112,"Kozani Prefecture",85,GR,Greece,58,,40.30055860,21.78877370
2106,Laconia,85,GR,Greece,16,,43.52785460,-71.47035090
2132,"Larissa Prefecture",85,GR,Greece,42,,39.63902240,22.41912540
2104,"Lefkada Regional Unit",85,GR,Greece,24,,38.83336630,20.70691080
2107,"Pella Regional Unit",85,GR,Greece,59,,40.91480390,22.14302150
2119,"Peloponnese Region",85,GR,Greece,J,,37.50794720,22.37349000
2114,"Phthiotis Prefecture",85,GR,Greece,06,,38.99978500,22.33377690
2103,"Preveza Prefecture",85,GR,Greece,34,,38.95926490,20.75171550
2121,"Serres Prefecture",85,GR,Greece,62,,41.08638540,23.54838190
2118,"South Aegean",85,GR,Greece,L,,37.08553020,25.14892150
2097,"Thessaloniki Regional Unit",85,GR,Greece,54,,40.64006290,22.94441910
2096,"West Greece Region",85,GR,Greece,G,,38.51154960,21.57067860
2108,"West Macedonia Region",85,GR,Greece,C,,40.30040580,21.79035590
3867,"Carriacou and Petite Martinique",87,GD,Grenada,10,,12.47858880,-61.44938420
3865,"Saint Andrew Parish",87,GD,Grenada,01,,,
3869,"Saint David Parish",87,GD,Grenada,02,,,
3864,"Saint George Parish",87,GD,Grenada,03,,,
3868,"Saint John Parish",87,GD,Grenada,04,,30.11183310,-90.48799160
3866,"Saint Mark Parish",87,GD,Grenada,05,,40.58818630,-73.94957010
3863,"Saint Patrick Parish",87,GD,Grenada,06,,,
3671,"Alta Verapaz Department",90,GT,Guatemala,AV,,15.59428830,-90.14949880
3674,"Baja Verapaz Department",90,GT,Guatemala,BV,,15.12558670,-90.37483540
3675,"Chimaltenango Department",90,GT,Guatemala,CM,,14.56347870,-90.98206680
3666,"Chiquimula Department",90,GT,Guatemala,CQ,,14.75149990,-89.47421770
3662,"El Progreso Department",90,GT,Guatemala,PR,,14.93887320,-90.07467670
3677,"Escuintla Department",90,GT,Guatemala,ES,,14.19109120,-90.98206680
3672,"Guatemala Department",90,GT,Guatemala,GU,,14.56494010,-90.52578230
3670,"Huehuetenango Department",90,GT,Guatemala,HU,,15.58799140,-91.67606910
3659,"Izabal Department",90,GT,Guatemala,IZ,,15.49765170,-88.86469800
3658,"Jalapa Department",90,GT,Guatemala,JA,,14.61214460,-89.96267990
3673,"Jutiapa Department",90,GT,Guatemala,JU,,14.19308020,-89.92532330
3669,"Petén Department",90,GT,Guatemala,PE,,16.91203300,-90.29957850
3668,"Quetzaltenango Department",90,GT,Guatemala,QZ,,14.79243300,-91.71495800
3657,"Quiché Department",90,GT,Guatemala,QC,,15.49838080,-90.98206680
3664,"Retalhuleu Department",90,GT,Guatemala,RE,,14.52454850,-91.68578800
3676,"Sacatepéquez Department",90,GT,Guatemala,SA,,14.51783790,-90.71527490
3667,"San Marcos Department",90,GT,Guatemala,SM,,14.93095690,-91.90992380
3665,"Santa Rosa Department",90,GT,Guatemala,SR,,38.44057590,-122.70375430
3661,"Sololá Department",90,GT,Guatemala,SO,,14.74852300,-91.28910360
3660,"Suchitepéquez Department",90,GT,Guatemala,SU,,14.42159820,-91.40482490
3663,"Totonicapán Department",90,GT,Guatemala,TO,,14.91734020,-91.36139230
2672,"Beyla Prefecture",92,GN,Guinea,BE,,8.91981780,-8.30884410
2699,"Boffa Prefecture",92,GN,Guinea,BF,,10.18082540,-14.03916150
2709,"Boké Prefecture",92,GN,Guinea,BK,,11.08473790,-14.37919120
2676,"Boké Region",92,GN,Guinea,B,,11.18646720,-14.10013260
2686,Conakry,92,GN,Guinea,C,,9.64118550,-13.57840120
2705,"Coyah Prefecture",92,GN,Guinea,CO,,9.77155350,-13.31252990
2679,"Dabola Prefecture",92,GN,Guinea,DB,,10.72978060,-11.11078540
2706,"Dalaba Prefecture",92,GN,Guinea,DL,,10.68681760,-12.24906970
2688,"Dinguiraye Prefecture",92,GN,Guinea,DI,,11.68442220,-10.80000510
2681,"Dubréka Prefecture",92,GN,Guinea,DU,,9.79073480,-13.51477350
2682,"Faranah Prefecture",92,GN,Guinea,FA,,9.90573990,-10.80000510
2683,"Forécariah Prefecture",92,GN,Guinea,FO,,9.38861870,-13.08179030
2675,"Fria Prefecture",92,GN,Guinea,FR,,10.36745430,-13.58418710
2685,"Gaoual Prefecture",92,GN,Guinea,GA,,11.57628040,-13.35872880
2711,"Guéckédou Prefecture",92,GN,Guinea,GU,,8.56496880,-10.13111630
2704,"Kankan Prefecture",92,GN,Guinea,KA,,10.30344650,-9.36730840
2697,"Kankan Region",92,GN,Guinea,K,,10.12092300,-9.54509740
2710,"Kérouané Prefecture",92,GN,Guinea,KE,,9.25366430,-9.01289260
2693,"Kindia Prefecture",92,GN,Guinea,KD,,10.10132920,-12.71351210
2701,"Kindia Region",92,GN,Guinea,D,,10.17816940,-12.98961500
2691,"Kissidougou Prefecture",92,GN,Guinea,KS,,9.22520220,-10.08072980
2692,"Koubia Prefecture",92,GN,Guinea,KB,,11.58235400,-11.89202370
2703,"Koundara Prefecture",92,GN,Guinea,KN,,12.48940210,-13.30675620
2695,"Kouroussa Prefecture",92,GN,Guinea,KO,,10.64892290,-9.88505860
2680,"Labé Prefecture",92,GN,Guinea,LA,,11.35419390,-12.34638750
2677,"Labé Region",92,GN,Guinea,L,,11.32320420,-12.28913140
2690,"Lélouma Prefecture",92,GN,Guinea,LE,,11.18333300,-12.93333300
2708,"Lola Prefecture",92,GN,Guinea,LO,,7.96138180,-8.39649380
2702,"Macenta Prefecture",92,GN,Guinea,MC,,8.46157950,-9.27855830
2700,"Mali Prefecture",92,GN,Guinea,ML,,11.98370900,-12.25479190
2689,"Mamou Prefecture",92,GN,Guinea,MM,,10.57360240,-11.88917210
2698,"Mamou Region",92,GN,Guinea,M,,10.57360240,-11.88917210
2673,"Mandiana Prefecture",92,GN,Guinea,MD,,10.61728270,-8.69857160
2678,"Nzérékoré Prefecture",92,GN,Guinea,NZ,,7.74783590,-8.82525020
2684,"Nzérékoré Region",92,GN,Guinea,N,,8.03858700,-8.83627550
2694,"Pita Prefecture",92,GN,Guinea,PI,,10.80620860,-12.71351210
2707,"Siguiri Prefecture",92,GN,Guinea,SI,,11.41481130,-9.17883040
2687,"Télimélé Prefecture",92,GN,Guinea,TE,,10.90893640,-13.02993310
2696,"Tougué Prefecture",92,GN,Guinea,TO,,11.38415830,-11.61577730
2674,"Yomou Prefecture",92,GN,Guinea,YO,,7.56962790,-9.25915710
2720,Bafatá,93,GW,Guinea-Bissau,BA,,12.17352430,-14.65295200
2714,"Biombo Region",93,GW,Guinea-Bissau,BM,,11.85290610,-15.73511710
2722,"Bolama Region",93,GW,Guinea-Bissau,BL,,11.14805910,-16.13457050
2713,"Cacheu Region",93,GW,Guinea-Bissau,CA,,12.05514160,-16.06401790
2719,"Gabú Region",93,GW,Guinea-Bissau,GA,,11.89624880,-14.10013260
2721,"Leste Province",93,GW,Guinea-Bissau,L,,,
2717,"Norte Province",93,GW,Guinea-Bissau,N,,7.87218110,123.88577470
2718,"Oio Region",93,GW,Guinea-Bissau,OI,,12.27607090,-15.31311850
2715,"Quinara Region",93,GW,Guinea-Bissau,QU,,11.79556200,-15.17268160
2716,"Sul Province",93,GW,Guinea-Bissau,S,,-10.28665780,20.71224650
2712,"Tombali Region",93,GW,Guinea-Bissau,TO,,11.36326960,-14.98561760
2764,Barima-Waini,94,GY,Guyana,BA,,7.48824190,-59.65644940
2760,Cuyuni-Mazaruni,94,GY,Guyana,CU,,6.46421410,-60.21107520
2767,Demerara-Mahaica,94,GY,Guyana,DE,,6.54642600,-58.09820460
2766,"East Berbice-Corentyne",94,GY,Guyana,EB,,2.74779220,-57.46272590
2768,"Essequibo Islands-West Demerara",94,GY,Guyana,ES,,6.57201320,-58.46299970
2762,Mahaica-Berbice,94,GY,Guyana,MA,,6.23849600,-57.91625550
2765,Pomeroon-Supenaam,94,GY,Guyana,PM,,7.12941660,-58.92062950
2761,Potaro-Siparuni,94,GY,Guyana,PT,,4.78558530,-59.28799770
2763,"Upper Demerara-Berbice",94,GY,Guyana,UD,,5.30648790,-58.18929210
2769,"Upper Takutu-Upper Essequibo",94,GY,Guyana,UT,,2.92395950,-58.73736340
4123,Artibonite,95,HT,Haiti,AR,,19.36290200,-72.42581450
4125,Centre,95,HT,Haiti,CE,,32.83702510,-96.77738820
4119,Grand'Anse,95,HT,Haiti,GA,,12.01666670,-61.76666670
4118,Nippes,95,HT,Haiti,NI,,18.39907350,-73.41802110
4117,Nord,95,HT,Haiti,ND,,43.19052600,-89.43792100
4121,Nord-Est,95,HT,Haiti,NE,,19.48897230,-71.85713310
4126,Nord-Ouest,95,HT,Haiti,NO,,19.83740090,-73.04052770
4120,Ouest,95,HT,Haiti,OU,,45.45472490,-73.65023650
4122,Sud,95,HT,Haiti,SD,,29.92132480,-90.09737720
4124,Sud-Est,95,HT,Haiti,SE,,18.27835980,-72.35479150
4047,"Atlántida Department",97,HN,Honduras,AT,,15.66962830,-87.14228950
4045,"Bay Islands Department",97,HN,Honduras,IB,,16.48266140,-85.87932520
4041,"Choluteca Department",97,HN,Honduras,CH,,13.25043250,-87.14228950
4051,"Colón Department",97,HN,Honduras,CL,,15.64259650,-85.52002400
4042,"Comayagua Department",97,HN,Honduras,CM,,14.55348280,-87.61863790
4049,"Copán Department",97,HN,Honduras,CP,,14.93608380,-88.86469800
4046,"Cortés Department",97,HN,Honduras,CR,,15.49235080,-88.09007620
4043,"El Paraíso Department",97,HN,Honduras,EP,,13.98212940,-86.49965460
4052,"Francisco Morazán Department",97,HN,Honduras,FM,,14.45411000,-87.06242610
4048,"Gracias a Dios Department",97,HN,Honduras,GD,,15.34180600,-84.60604490
4044,"Intibucá Department",97,HN,Honduras,IN,,14.37273400,-88.24611830
4058,"La Paz Department",97,HN,Honduras,LP,,-15.08924160,-68.52471490
4054,"Lempira Department",97,HN,Honduras,LE,,14.18876980,-88.55653100
4056,"Ocotepeque Department",97,HN,Honduras,OC,,14.51703470,-89.05615320
4050,"Olancho Department",97,HN,Honduras,OL,,14.80674060,-85.76666450
4053,"Santa Bárbara Department",97,HN,Honduras,SB,,15.12027950,-88.40160410
4055,"Valle Department",97,HN,Honduras,VA,,13.57829360,-87.57912870
4057,"Yoro Department",97,HN,Honduras,YO,,15.29496790,-87.14228950
4889,"Central and Western",98,HK,"Hong Kong S.A.R.",HCW,district,22.28666000,114.15497000
4891,Eastern,98,HK,"Hong Kong S.A.R.",HEA,district,22.28411000,114.22414000
4888,Islands,98,HK,"Hong Kong S.A.R.",NIS,district,22.26114000,113.94608000
4895,"Kowloon City",98,HK,"Hong Kong S.A.R.",KKC,district,22.32820000,114.19155000
4898,"Kwai Tsing",98,HK,"Hong Kong S.A.R.",NKT,district,22.35488000,114.08401000
4897,"Kwun Tong",98,HK,"Hong Kong S.A.R.",KKT,district,22.31326000,114.22581000
4900,North,98,HK,"Hong Kong S.A.R.",NNO,district,22.49471000,114.13812000
4887,"Sai Kung",98,HK,"Hong Kong S.A.R.",NSK,district,22.38143000,114.27052000
4901,"Sha Tin",98,HK,"Hong Kong S.A.R.",NST,district,22.38715000,114.19534000
4894,"Sham Shui Po",98,HK,"Hong Kong S.A.R.",KSS,district,22.33074000,114.16220000
4892,Southern,98,HK,"Hong Kong S.A.R.",HSO,district,22.24725000,114.15884000
4885,"Tai Po",98,HK,"Hong Kong S.A.R.",NTP,district,22.45085000,114.16422000
4884,"Tsuen Wan",98,HK,"Hong Kong S.A.R.",NTW,district,22.36281000,114.12907000
4899,"Tuen Mun",98,HK,"Hong Kong S.A.R.",NTM,district,22.39163000,113.97708850
4890,"Wan Chai",98,HK,"Hong Kong S.A.R.",HWC,district,22.27968000,114.17168000
4896,"Wong Tai Sin",98,HK,"Hong Kong S.A.R.",KWT,district,22.33353000,114.19686000
4893,"Yau Tsim Mong",98,HK,"Hong Kong S.A.R.",KYT,district,22.32138000,114.17260000
4883,"Yuen Long",98,HK,"Hong Kong S.A.R.",NYL,district,22.44559000,114.02218000
1048,Bács-Kiskun,99,HU,Hungary,BK,county,46.56614370,19.42724640
1055,Baranya,99,HU,Hungary,BA,county,46.04845850,18.27191730
1060,Békés,99,HU,Hungary,BE,county,46.67048990,21.04349960
1036,Békéscsaba,99,HU,Hungary,BC,"city with county rights",46.67359390,21.08773090
1058,Borsod-Abaúj-Zemplén,99,HU,Hungary,BZ,county,48.29394010,20.69341120
1064,Budapest,99,HU,Hungary,BU,"capital city",47.49791200,19.04023500
1031,"Csongrád County",99,HU,Hungary,CS,county,46.41670500,20.25661610
1032,Debrecen,99,HU,Hungary,DE,"city with county rights",47.53160490,21.62731240
1049,Dunaújváros,99,HU,Hungary,DU,"city with county rights",46.96190590,18.93552270
1037,Eger,99,HU,Hungary,EG,"city with county rights",47.90253480,20.37722840
1028,Érd,99,HU,Hungary,ER,"city with county rights",47.39197180,18.90454400
1044,"Fejér County",99,HU,Hungary,FE,county,47.12179320,18.52948150
1041,Győr,99,HU,Hungary,GY,"city with county rights",47.68745690,17.65039740
1042,"Győr-Moson-Sopron County",99,HU,Hungary,GS,county,47.65092850,17.25058830
1063,"Hajdú-Bihar County",99,HU,Hungary,HB,county,47.46883550,21.54532270
1040,"Heves County",99,HU,Hungary,HE,county,47.80576170,20.20385590
1027,Hódmezővásárhely,99,HU,Hungary,HV,"city with county rights",46.41812620,20.33003150
1043,"Jász-Nagykun-Szolnok County",99,HU,Hungary,JN,county,47.25555790,20.52324560
1067,Kaposvár,99,HU,Hungary,KV,"city with county rights",46.35936060,17.79676390
1056,Kecskemét,99,HU,Hungary,KM,"city with county rights",46.89637110,19.68968610
5085,Komárom-Esztergom,99,HU,Hungary,KE,county,47.57797860,18.12568550
1065,Miskolc,99,HU,Hungary,MI,"city with county rights",48.10347750,20.77843840
1030,Nagykanizsa,99,HU,Hungary,NK,"city with county rights",46.45902180,16.98967960
1051,"Nógrád County",99,HU,Hungary,NO,county,47.90410310,19.04985040
1034,Nyíregyháza,99,HU,Hungary,NY,"city with county rights",47.94953240,21.72440530
1053,Pécs,99,HU,Hungary,PS,"city with county rights",46.07273450,18.23226600
1059,"Pest County",99,HU,Hungary,PE,county,47.44800010,19.46181280
1068,Salgótarján,99,HU,Hungary,ST,"city with county rights",48.09352370,19.79998130
1035,"Somogy County",99,HU,Hungary,SO,county,46.55485900,17.58667320
1057,Sopron,99,HU,Hungary,SN,"city with county rights",47.68166190,16.58447950
1045,"Szabolcs-Szatmár-Bereg County",99,HU,Hungary,SZ,county,48.03949540,22.00333000
1029,Szeged,99,HU,Hungary,SD,"city with county rights",46.25301020,20.14142530
1033,Székesfehérvár,99,HU,Hungary,SF,"city with county rights",47.18602620,18.42213580
1061,Szekszárd,99,HU,Hungary,SS,"city with county rights",46.34743260,18.70622930
1047,Szolnok,99,HU,Hungary,SK,"city with county rights",47.16213550,20.18247120
1052,Szombathely,99,HU,Hungary,SH,"city with county rights",47.23068510,16.62184410
1066,Tatabánya,99,HU,Hungary,TB,"city with county rights",47.56924600,18.40481800
1038,"Tolna County",99,HU,Hungary,TO,county,46.47627540,18.55706270
1039,"Vas County",99,HU,Hungary,VA,county,47.09291110,16.68121830
1062,Veszprém,99,HU,Hungary,VM,"city with county rights",47.10280870,17.90930190
1054,"Veszprém County",99,HU,Hungary,VE,county,47.09309740,17.91007630
1046,"Zala County",99,HU,Hungary,ZA,county,46.73844040,16.91522520
1050,Zalaegerszeg,99,HU,Hungary,ZE,county,46.84169360,16.84163220
3431,"Capital Region",100,IS,Iceland,1,,38.56569570,-92.18169490
3433,"Eastern Region",100,IS,Iceland,7,,,
3437,"Northeastern Region",100,IS,Iceland,6,,43.29942850,-74.21793260
3435,"Northwestern Region",100,IS,Iceland,5,,41.91339320,-73.04716880
3430,"Southern Peninsula Region",100,IS,Iceland,2,,63.91548030,-22.36496670
3434,"Southern Region",100,IS,Iceland,8,,,
3436,"Western Region",100,IS,Iceland,3,,,
3432,Westfjords,100,IS,Iceland,4,,65.91961500,-21.88117640
4023,"Andaman and Nicobar Islands",101,IN,India,AN,"Union territory",11.74008670,92.65864010
4017,"Andhra Pradesh",101,IN,India,AP,state,15.91289980,79.73998750
4024,"Arunachal Pradesh",101,IN,India,AR,state,28.21799940,94.72775280
4027,Assam,101,IN,India,AS,state,26.20060430,92.93757390
4037,Bihar,101,IN,India,BR,state,25.09607420,85.31311940
4031,Chandigarh,101,IN,India,CH,"Union territory",30.73331480,76.77941790
4040,Chhattisgarh,101,IN,India,CT,state,21.27865670,81.86614420
4033,"Dadra and Nagar Haveli and Daman and Diu",101,IN,India,DH,"Union territory",20.39737360,72.83279910
4021,Delhi,101,IN,India,DL,"Union territory",28.70405920,77.10249020
4009,Goa,101,IN,India,GA,state,15.29932650,74.12399600
4030,Gujarat,101,IN,India,GJ,state,22.25865200,71.19238050
4007,Haryana,101,IN,India,HR,state,29.05877570,76.08560100
4020,"Himachal Pradesh",101,IN,India,HP,state,31.10482940,77.17339010
4029,"Jammu and Kashmir",101,IN,India,JK,"Union territory",33.27783900,75.34121790
4025,Jharkhand,101,IN,India,JH,state,23.61018080,85.27993540
4026,Karnataka,101,IN,India,KA,state,15.31727750,75.71388840
4028,Kerala,101,IN,India,KL,state,10.85051590,76.27108330
4852,Ladakh,101,IN,India,LA,"Union territory",34.22684750,77.56194190
4019,Lakshadweep,101,IN,India,LD,"Union territory",10.32802650,72.78463360
4039,"Madhya Pradesh",101,IN,India,MP,state,22.97342290,78.65689420
4008,Maharashtra,101,IN,India,MH,state,19.75147980,75.71388840
4010,Manipur,101,IN,India,MN,state,24.66371730,93.90626880
4006,Meghalaya,101,IN,India,ML,state,25.46703080,91.36621600
4036,Mizoram,101,IN,India,MZ,state,23.16454300,92.93757390
4018,Nagaland,101,IN,India,NL,state,26.15843540,94.56244260
4013,Odisha,101,IN,India,OR,state,20.95166580,85.09852360
4011,Puducherry,101,IN,India,PY,"Union territory",11.94159150,79.80831330
4015,Punjab,101,IN,India,PB,state,31.14713050,75.34121790
4014,Rajasthan,101,IN,India,RJ,state,27.02380360,74.21793260
4034,Sikkim,101,IN,India,SK,state,27.53297180,88.51221780
4035,"Tamil Nadu",101,IN,India,TN,state,11.12712250,78.65689420
4012,Telangana,101,IN,India,TG,state,18.11243720,79.01929970
4038,Tripura,101,IN,India,TR,state,23.94084820,91.98815270
4022,"Uttar Pradesh",101,IN,India,UP,state,26.84670880,80.94615920
4016,Uttarakhand,101,IN,India,UK,state,30.06675300,79.01929970
4853,"West Bengal",101,IN,India,WB,state,22.98675690,87.85497550
1822,Aceh,102,ID,Indonesia,AC,province,4.69513500,96.74939930
1826,Bali,102,ID,Indonesia,BA,province,-8.34053890,115.09195090
1810,Banten,102,ID,Indonesia,BT,province,-6.40581720,106.06401790
1793,Bengkulu,102,ID,Indonesia,BE,province,-3.79284510,102.26076410
1829,"DI Yogyakarta",102,ID,Indonesia,YO,province,-7.87538490,110.42620880
1805,"DKI Jakarta",102,ID,Indonesia,JK,"capital district",-6.20876340,106.84559900
1812,Gorontalo,102,ID,Indonesia,GO,province,0.54354420,123.05676930
1815,Jambi,102,ID,Indonesia,JA,province,-1.61012290,103.61312030
1825,"Jawa Barat",102,ID,Indonesia,JB,province,-7.09091100,107.66888700
1802,"Jawa Tengah",102,ID,Indonesia,JT,province,-7.15097500,110.14025940
1827,"Jawa Timur",102,ID,Indonesia,JI,province,-7.53606390,112.23840170
1806,"Kalimantan Barat",102,ID,Indonesia,KB,province,0.47734750,106.61314050
1819,"Kalimantan Selatan",102,ID,Indonesia,KS,province,-3.09264150,115.28375850
1794,"Kalimantan Tengah",102,ID,Indonesia,KT,province,-1.68148780,113.38235450
1804,"Kalimantan Timur",102,ID,Indonesia,KI,province,0.53865860,116.41938900
1824,"Kalimantan Utara",102,ID,Indonesia,KU,province,3.07309290,116.04138890
1820,"Kepulauan Bangka Belitung",102,ID,Indonesia,BB,province,-2.74105130,106.44058720
1807,"Kepulauan Riau",102,ID,Indonesia,KR,province,3.94565140,108.14286690
1811,Lampung,102,ID,Indonesia,LA,province,-4.55858490,105.40680790
1800,Maluku,102,ID,Indonesia,MA,province,-3.23846160,130.14527340
1801,"Maluku Utara",102,ID,Indonesia,MU,province,1.57099930,127.80876930
1814,"Nusa Tenggara Barat",102,ID,Indonesia,NB,province,-8.65293340,117.36164760
1818,"Nusa Tenggara Timur",102,ID,Indonesia,NT,province,-8.65738190,121.07937050
1798,Papua,102,ID,Indonesia,PA,province,-5.01222020,141.34701590
1799,"Papua Barat",102,ID,Indonesia,PB,province,-1.33611540,133.17471620
1809,Riau,102,ID,Indonesia,RI,province,0.29334690,101.70682940
1817,"Sulawesi Barat",102,ID,Indonesia,SR,province,-2.84413710,119.23207840
1795,"Sulawesi Selatan",102,ID,Indonesia,SN,province,-3.66879940,119.97405340
1813,"Sulawesi Tengah",102,ID,Indonesia,ST,province,-1.43002540,121.44561790
1796,"Sulawesi Tenggara",102,ID,Indonesia,SG,province,-4.14491000,122.17460500
1808,"Sulawesi Utara",102,ID,Indonesia,SA,province,0.62469320,123.97500180
1828,"Sumatera Barat",102,ID,Indonesia,SB,province,-0.73993970,100.80000510
1816,"Sumatera Selatan",102,ID,Indonesia,SS,province,-3.31943740,103.91439900
1792,"Sumatera Utara",102,ID,Indonesia,SU,province,2.11535470,99.54509740
3929,Alborz,103,IR,Iran,30,province,35.99604670,50.92892460
3934,Ardabil,103,IR,Iran,24,province,38.48532760,47.89112090
3932,Bushehr,103,IR,Iran,18,province,28.76207390,51.51500770
3921,"Chaharmahal and Bakhtiari",103,IR,Iran,14,province,31.99704190,50.66138490
3944,"East Azerbaijan",103,IR,Iran,03,province,37.90357330,46.26821090
3939,Fars,103,IR,Iran,07,province,29.10438130,53.04589300
3920,Gilan,103,IR,Iran,01,province,37.28094550,49.59241340
3933,Golestan,103,IR,Iran,27,province,37.28981230,55.13758340
4920,Hamadan,103,IR,Iran,13,province,34.91936070,47.48329250
3937,Hormozgan,103,IR,Iran,22,province,27.13872300,55.13758340
3918,Ilam,103,IR,Iran,16,province,33.29576180,46.67053400
3923,Isfahan,103,IR,Iran,10,province,33.27710730,52.36133780
3943,Kerman,103,IR,Iran,08,province,29.48500890,57.64390480
3919,Kermanshah,103,IR,Iran,05,province,34.45762330,46.67053400
3917,Khuzestan,103,IR,Iran,06,province,31.43601490,49.04131200
3926,"Kohgiluyeh and Boyer-Ahmad",103,IR,Iran,17,province,30.72458600,50.84563230
3935,Kurdistan,103,IR,Iran,12,province,35.95535790,47.13621250
3928,Lorestan,103,IR,Iran,15,province,33.58183940,48.39881860
3916,Markazi,103,IR,Iran,00,province,34.61230500,49.85472660
3938,Mazandaran,103,IR,Iran,02,province,36.22623930,52.53186040
3942,"North Khorasan",103,IR,Iran,28,province,37.47103530,57.10131880
3941,Qazvin,103,IR,Iran,26,province,36.08813170,49.85472660
3922,Qom,103,IR,Iran,25,province,34.64157640,50.87460350
3927,"Razavi Khorasan",103,IR,Iran,09,province,35.10202530,59.10417580
3940,Semnan,103,IR,Iran,20,province,35.22555850,54.43421380
3931,"Sistan and Baluchestan",103,IR,Iran,11,province,27.52999060,60.58206760
3930,"South Khorasan",103,IR,Iran,29,province,32.51756430,59.10417580
3945,Tehran,103,IR,Iran,23,province,35.72484160,51.38165300
3924,"West Azarbaijan",103,IR,Iran,04,province,37.45500620,45.00000000
3936,Yazd,103,IR,Iran,21,province,32.10063870,54.43421380
3925,Zanjan,103,IR,Iran,19,province,36.50181850,48.39881860
3964,"Al Anbar",104,IQ,Iraq,AN,governorate,32.55976140,41.91964710
3958,"Al Muthanna",104,IQ,Iraq,MU,governorate,29.91331710,45.29938620
3956,Al-Qādisiyyah,104,IQ,Iraq,QA,governorate,32.04369100,45.14945050
3955,Babylon,104,IQ,Iraq,BB,governorate,32.46819100,44.55019350
3959,Baghdad,104,IQ,Iraq,BG,governorate,33.31526180,44.36606530
3960,Basra,104,IQ,Iraq,BA,governorate,30.51142520,47.82962530
3954,"Dhi Qar",104,IQ,Iraq,DQ,governorate,31.10422920,46.36246860
3965,Diyala,104,IQ,Iraq,DI,governorate,33.77334870,45.14945050
3967,Dohuk,104,IQ,Iraq,DA,governorate,36.90772520,43.06316890
3968,Erbil,104,IQ,Iraq,AR,governorate,36.55706280,44.38512630
3957,Karbala,104,IQ,Iraq,KA,governorate,32.40454930,43.86732220
3971,Kirkuk,104,IQ,Iraq,KI,governorate,35.32920140,43.94367880
3966,Maysan,104,IQ,Iraq,MA,governorate,31.87340020,47.13621250
3962,Najaf,104,IQ,Iraq,NA,governorate,31.35174860,44.09603110
3963,Nineveh,104,IQ,Iraq,NI,governorate,36.22957400,42.23624350
3961,Saladin,104,IQ,Iraq,SD,governorate,34.53375270,43.48373800
3969,Sulaymaniyah,104,IQ,Iraq,SU,governorate,35.54663480,45.30036830
3970,Wasit,104,IQ,Iraq,WA,governorate,32.60240940,45.75209850
1095,Carlow,105,IE,Ireland,CW,county,52.72322170,-6.81082950
1088,Cavan,105,IE,Ireland,CN,county,53.97654240,-7.29966230
1091,Clare,105,IE,Ireland,CE,county,43.04664000,-87.89958100
1087,Connacht,105,IE,Ireland,C,province,53.83762430,-8.95844810
1074,Cork,105,IE,Ireland,CO,county,51.89851430,-8.47560350
1071,Donegal,105,IE,Ireland,DL,county,54.65489930,-8.10409670
1072,Dublin,105,IE,Ireland,D,county,53.34980530,-6.26030970
1079,Galway,105,IE,Ireland,G,county,53.35645090,-8.85341130
1077,Kerry,105,IE,Ireland,KY,county,52.15446070,-9.56686330
1082,Kildare,105,IE,Ireland,KE,county,53.21204340,-6.81947080
1090,Kilkenny,105,IE,Ireland,KK,county,52.57769570,-7.21800200
1096,Laois,105,IE,Ireland,LS,county,52.99429500,-7.33230070
1073,Leinster,105,IE,Ireland,L,province,53.32715380,-7.51408410
1094,Limerick,105,IE,Ireland,LK,county,52.50905170,-8.74749550
1076,Longford,105,IE,Ireland,LD,county,53.72749820,-7.79315270
1083,Louth,105,IE,Ireland,LH,county,53.92523240,-6.48894230
1084,Mayo,105,IE,Ireland,MO,county,54.01526040,-9.42893690
1092,Meath,105,IE,Ireland,MH,county,53.60554800,-6.65641690
1075,Monaghan,105,IE,Ireland,MN,county,54.24920460,-6.96831320
1080,Munster,105,IE,Ireland,M,province,51.94711970,7.58453200
1078,Offaly,105,IE,Ireland,OY,county,53.23568710,-7.71222290
1081,Roscommon,105,IE,Ireland,RN,county,53.75926040,-8.26816210
1070,Sligo,105,IE,Ireland,SO,county,54.15532770,-8.60645320
1069,Tipperary,105,IE,Ireland,TA,county,52.47378940,-8.16185140
1086,Ulster,105,IE,Ireland,U,province,54.76165550,-6.96122730
1089,Waterford,105,IE,Ireland,WD,county,52.19435490,-7.62275120
1097,Westmeath,105,IE,Ireland,WH,county,53.53453080,-7.46532170
1093,Wexford,105,IE,Ireland,WX,county,52.47936030,-6.58399130
1085,Wicklow,105,IE,Ireland,WW,county,52.98623130,-6.36725430
1367,"Central District",106,IL,Israel,M,,47.60875830,-122.29642350
1369,"Haifa District",106,IL,Israel,HA,,32.48141110,34.99475100
1370,"Jerusalem District",106,IL,Israel,JM,,31.76482430,34.99475100
1366,"Northern District",106,IL,Israel,Z,,36.15118640,-95.99517630
1368,"Southern District",106,IL,Israel,D,,40.71375860,-74.00090590
1371,"Tel Aviv District",106,IL,Israel,TA,,32.09290750,34.80721650
1679,Abruzzo,107,IT,Italy,65,region,42.19201190,13.72891670
1727,Agrigento,107,IT,Italy,AG,"free municipal consortium",37.31052020,13.58579780
1783,Alessandria,107,IT,Italy,AL,province,44.81755870,8.70466270
1672,Ancona,107,IT,Italy,AN,province,43.54932450,13.26634790
1716,"Aosta Valley",107,IT,Italy,23,"autonomous region",45.73888780,7.42618660
1688,Apulia,107,IT,Italy,75,region,40.79283930,17.10119310
1681,"Ascoli Piceno",107,IT,Italy,AP,province,42.86389330,13.58997330
1780,Asti,107,IT,Italy,AT,province,44.90076520,8.20643150
1692,Avellino,107,IT,Italy,AV,province,40.99645100,15.12589550
1686,Barletta-Andria-Trani,107,IT,Italy,BT,province,41.20045430,16.20514840
1706,Basilicata,107,IT,Italy,77,region,40.64307660,15.96998780
1689,Belluno,107,IT,Italy,BL,province,46.24976590,12.19695650
1701,Benevento,107,IT,Italy,BN,province,41.20350930,14.75209390
1704,Bergamo,107,IT,Italy,BG,province,45.69826420,9.67726980
1778,Biella,107,IT,Italy,BI,province,45.56281760,8.05827170
1717,Brescia,107,IT,Italy,BS,province,45.54155260,10.21180190
1714,Brindisi,107,IT,Italy,BR,province,40.61126630,17.76362100
1703,Calabria,107,IT,Italy,78,region,39.30877140,16.34637910
1718,Caltanissetta,107,IT,Italy,CL,"free municipal consortium",37.48601300,14.06149820
1669,Campania,107,IT,Italy,72,region,40.66708870,15.10681130
1721,Campobasso,107,IT,Italy,CB,province,41.67388650,14.75209390
1731,Caserta,107,IT,Italy,CE,province,41.20783540,14.10013260
1728,Catanzaro,107,IT,Italy,CZ,province,38.88963480,16.44058720
1739,Chieti,107,IT,Italy,CH,province,42.03344280,14.37919120
1740,Como,107,IT,Italy,CO,province,45.80804160,9.08517930
1742,Cosenza,107,IT,Italy,CS,province,39.56441050,16.25221430
1751,Cremona,107,IT,Italy,CR,province,45.20143750,9.98365820
1754,Crotone,107,IT,Italy,KR,province,39.13098560,17.00670310
1775,Cuneo,107,IT,Italy,CN,province,44.59703140,7.61142170
1773,Emilia-Romagna,107,IT,Italy,45,region,44.59676070,11.21863960
1723,Enna,107,IT,Italy,EN,"free municipal consortium",37.56762160,14.27953490
1744,Fermo,107,IT,Italy,FM,province,43.09313670,13.58997330
1746,Ferrara,107,IT,Italy,FE,province,44.76636800,11.76440680
1771,Foggia,107,IT,Italy,FG,province,41.63844800,15.59433880
1779,Forlì-Cesena,107,IT,Italy,FC,province,43.99476810,11.98046130
1756,"Friuli–Venezia Giulia",107,IT,Italy,36,"autonomous region",46.22591770,13.10336460
1776,Frosinone,107,IT,Italy,FR,province,41.65765280,13.63627150
1777,Gorizia,107,IT,Italy,GO,"decentralized regional entity",45.90538990,13.51637250
1787,Grosseto,107,IT,Italy,GR,province,42.85180070,11.25237920
1788,Imperia,107,IT,Italy,IM,province,43.94186600,7.82863680
1789,Isernia,107,IT,Italy,IS,province,41.58915550,14.19309180
1781,L'Aquila,107,IT,Italy,AQ,province,42.12563170,13.63627150
1791,"La Spezia",107,IT,Italy,SP,province,44.24479130,9.76786870
1674,Latina,107,IT,Italy,LT,province,41.40874760,13.08179030
1678,Lazio,107,IT,Italy,62,region,41.81224100,12.73851000
1675,Lecce,107,IT,Italy,LE,province,40.23473930,18.14286690
1677,Lecco,107,IT,Italy,LC,province,45.93829410,9.38572900
1768,Liguria,107,IT,Italy,42,region,44.31679170,8.39649380
1745,Livorno,107,IT,Italy,LI,province,43.02398480,10.66471010
1747,Lodi,107,IT,Italy,LO,province,45.24050360,9.52925120
1705,Lombardy,107,IT,Italy,25,region,45.47906710,9.84524330
1749,Lucca,107,IT,Italy,LU,province,43.83767360,10.49505300
1750,Macerata,107,IT,Italy,MC,province,43.24593220,13.26634790
1758,Mantua,107,IT,Italy,MN,province,45.16677280,10.77536130
1670,Marche,107,IT,Italy,57,region,43.30456200,13.71947000
1759,"Massa and Carrara",107,IT,Italy,MS,province,44.22139980,10.03596610
1760,Matera,107,IT,Italy,MT,province,40.66634960,16.60436360
1761,"Medio Campidano",107,IT,Italy,VS,province,39.53173890,8.70407500
1757,Modena,107,IT,Italy,MO,province,44.55137990,10.91804800
1695,Molise,107,IT,Italy,67,region,41.67388650,14.75209390
1769,"Monza and Brianza",107,IT,Italy,MB,province,45.62359900,9.25880150
1774,Novara,107,IT,Italy,NO,province,45.54851330,8.51507930
1790,Nuoro,107,IT,Italy,NU,province,40.32869040,9.45615500
1786,Oristano,107,IT,Italy,OR,province,40.05990680,8.74811670
1665,Padua,107,IT,Italy,PD,province,45.36618640,11.82091390
1668,Palermo,107,IT,Italy,PA,province,38.11569000,13.36148680
1666,Parma,107,IT,Italy,PR,province,44.80153220,10.32793540
1676,Pavia,107,IT,Italy,PV,province,45.32181660,8.84662360
1691,Perugia,107,IT,Italy,PG,province,42.93800400,12.62162110
1693,"Pesaro and Urbino",107,IT,Italy,PU,province,43.61301180,12.71351210
1694,Pescara,107,IT,Italy,PE,province,42.35706550,13.96080910
1696,Piacenza,107,IT,Italy,PC,province,44.82631120,9.52914470
1702,Piedmont,107,IT,Italy,21,region,45.05223660,7.51538850
1685,Pisa,107,IT,Italy,PI,province,43.72283150,10.40171940
1687,Pistoia,107,IT,Italy,PT,province,43.95437330,10.89030990
1690,Pordenone,107,IT,Italy,PN,"decentralized regional entity",46.03788620,12.71083500
1697,Potenza,107,IT,Italy,PZ,province,40.41821940,15.87600400
1700,Prato,107,IT,Italy,PO,province,44.04539000,11.11644520
1729,Ragusa,107,IT,Italy,RG,"free municipal consortium",36.92692730,14.72551290
1707,Ravenna,107,IT,Italy,RA,province,44.41844430,12.20359980
1708,"Reggio Emilia",107,IT,Italy,RE,province,44.58565800,10.55647360
1712,Rieti,107,IT,Italy,RI,province,42.36744050,12.89750980
1713,Rimini,107,IT,Italy,RN,province,44.06782880,12.56951580
1719,Rovigo,107,IT,Italy,RO,province,45.02418180,11.82381620
1720,Salerno,107,IT,Italy,SA,province,40.42878320,15.21948080
1715,Sardinia,107,IT,Italy,88,"autonomous region",40.12087520,9.01289260
1722,Sassari,107,IT,Italy,SS,province,40.79679070,8.57504070
1732,Savona,107,IT,Italy,SV,province,44.28879950,8.26505800
1709,Sicily,107,IT,Italy,82,"autonomous region",37.59999380,14.01535570
1734,Siena,107,IT,Italy,SI,province,43.29377320,11.43391480
1667,Siracusa,107,IT,Italy,SR,"free municipal consortium",37.06569240,15.28571090
1741,Sondrio,107,IT,Italy,SO,province,46.17276360,9.79949170
1730,"South Sardinia",107,IT,Italy,SU,province,39.38935350,8.93970000
1743,Taranto,107,IT,Italy,TA,province,40.57409010,17.24299760
1752,Teramo,107,IT,Italy,TE,province,42.58956080,13.63627150
1755,Terni,107,IT,Italy,TR,province,42.56345340,12.52980280
1733,Trapani,107,IT,Italy,TP,"free municipal consortium",38.01831160,12.51482650
1725,"Trentino-South Tyrol",107,IT,Italy,32,"autonomous region",46.43366620,11.16932960
1762,Treviso,107,IT,Italy,TV,province,45.66685170,12.24306170
1763,Trieste,107,IT,Italy,TS,"decentralized regional entity",45.68948230,13.78330720
1664,Tuscany,107,IT,Italy,52,region,43.77105130,11.24862080
1764,Udine,107,IT,Italy,UD,"decentralized regional entity",46.14079720,13.16628960
1683,Umbria,107,IT,Italy,55,region,42.93800400,12.62162110
1765,Varese,107,IT,Italy,VA,province,45.79902600,8.73009450
1753,Veneto,107,IT,Italy,34,region,45.44146620,12.31525950
1726,Verbano-Cusio-Ossola,107,IT,Italy,VB,province,46.13996880,8.27246490
1785,Vercelli,107,IT,Italy,VC,province,45.32022040,8.41850800
1736,Verona,107,IT,Italy,VR,province,45.44184980,11.07353160
1737,"Vibo Valentia",107,IT,Italy,VV,province,38.63785650,16.20514840
1738,Vicenza,107,IT,Italy,VI,province,45.69839950,11.56614650
1735,Viterbo,107,IT,Italy,VT,province,42.40024200,11.88917210
3753,"Clarendon Parish",108,JM,Jamaica,13,,17.95571830,-77.24051530
3749,"Hanover Parish",108,JM,Jamaica,09,,18.40977070,-78.13363800
3748,"Kingston Parish",108,JM,Jamaica,01,,17.96832710,-76.78270200
3754,"Manchester Parish",108,JM,Jamaica,12,,18.06696540,-77.51607880
3752,"Portland Parish",108,JM,Jamaica,04,,18.08442740,-76.41002670
3751,"Saint Andrew",108,JM,Jamaica,02,,37.22451030,-95.70211890
3744,"Saint Ann Parish",108,JM,Jamaica,06,,37.28714520,-77.41035330
3746,"Saint Catherine Parish",108,JM,Jamaica,14,,18.03641340,-77.05644640
3743,"Saint Elizabeth Parish",108,JM,Jamaica,11,,38.99253080,-94.58992000
3745,"Saint James Parish",108,JM,Jamaica,08,,30.01792920,-90.79132270
3747,"Saint Mary Parish",108,JM,Jamaica,05,,36.09252200,-95.97384400
3750,"Saint Thomas Parish",108,JM,Jamaica,03,,41.44253890,-81.74022180
3755,"Trelawny Parish",108,JM,Jamaica,07,,18.35261430,-77.60778650
3742,"Westmoreland Parish",108,JM,Jamaica,10,,18.29443780,-78.15644320
827,"Aichi Prefecture",109,JP,Japan,23,,35.01825050,137.29238930
829,"Akita Prefecture",109,JP,Japan,05,,40.13762930,140.33434100
839,"Aomori Prefecture",109,JP,Japan,02,,40.76570770,140.91758790
821,"Chiba Prefecture",109,JP,Japan,12,,35.33541550,140.18325160
865,"Ehime Prefecture",109,JP,Japan,38,,33.60253060,132.78575830
848,"Fukui Prefecture",109,JP,Japan,18,,35.89622700,136.21115790
861,"Fukuoka Prefecture",109,JP,Japan,40,,33.56625590,130.71585700
847,"Fukushima Prefecture",109,JP,Japan,07,,37.38343730,140.18325160
858,"Gifu Prefecture",109,JP,Japan,21,,35.74374910,136.98051030
862,"Gunma Prefecture",109,JP,Japan,10,,36.56053880,138.87999720
828,"Hiroshima Prefecture",109,JP,Japan,34,,34.88234080,133.01948970
832,"Hokkaidō Prefecture",109,JP,Japan,01,,43.22032660,142.86347370
831,"Hyōgo Prefecture",109,JP,Japan,28,,34.85795180,134.54537870
851,"Ibaraki Prefecture",109,JP,Japan,08,,36.21935710,140.18325160
830,"Ishikawa Prefecture",109,JP,Japan,17,,36.32603170,136.52896530
856,"Iwate Prefecture",109,JP,Japan,03,,39.58329890,141.25345740
864,"Kagawa Prefecture",109,JP,Japan,37,,34.22259150,134.01991520
840,"Kagoshima Prefecture",109,JP,Japan,46,,31.39119580,130.87785860
842,"Kanagawa Prefecture",109,JP,Japan,14,,35.49135350,139.28414300
4924,"Kōchi Prefecture",109,JP,Japan,39,,33.28791610,132.27592620
846,"Kumamoto Prefecture",109,JP,Japan,43,,32.85944270,130.79691490
834,"Kyōto Prefecture",109,JP,Japan,26,,35.15666090,135.52519820
833,"Mie Prefecture",109,JP,Japan,24,,33.81439010,136.04870470
857,"Miyagi Prefecture",109,JP,Japan,04,,38.63061200,141.11930480
855,"Miyazaki Prefecture",109,JP,Japan,45,,32.60360220,131.44125100
843,"Nagano Prefecture",109,JP,Japan,20,,36.15439410,137.92182040
849,"Nagasaki Prefecture",109,JP,Japan,42,,33.24885250,129.69309120
824,"Nara Prefecture",109,JP,Japan,29,,34.29755280,135.82797340
841,"Niigata Prefecture",109,JP,Japan,15,,37.51783860,138.92697940
822,"Ōita Prefecture",109,JP,Japan,44,,33.15892990,131.36111210
820,"Okayama Prefecture",109,JP,Japan,33,,34.89634070,133.63753140
853,"Okinawa Prefecture",109,JP,Japan,47,,26.12019110,127.70250120
859,"Ōsaka Prefecture",109,JP,Japan,27,,34.64133150,135.56293940
863,"Saga Prefecture",109,JP,Japan,41,,33.30783710,130.22712430
860,"Saitama Prefecture",109,JP,Japan,11,,35.99625130,139.44660050
845,"Shiga Prefecture",109,JP,Japan,25,,35.32920140,136.05632120
826,"Shimane Prefecture",109,JP,Japan,32,,35.12440940,132.62934460
825,"Shizuoka Prefecture",109,JP,Japan,22,,35.09293970,138.31902760
854,"Tochigi Prefecture",109,JP,Japan,09,,36.67147390,139.85472660
836,"Tokushima Prefecture",109,JP,Japan,36,,33.94196550,134.32365570
823,Tokyo,109,JP,Japan,13,,35.67619190,139.65031060
850,"Tottori Prefecture",109,JP,Japan,31,,35.35731610,133.40666180
838,"Toyama Prefecture",109,JP,Japan,16,,36.69582660,137.21370710
844,"Wakayama Prefecture",109,JP,Japan,30,,33.94809140,135.37453580
837,"Yamagata Prefecture",109,JP,Japan,06,,38.53705640,140.14351980
835,"Yamaguchi Prefecture",109,JP,Japan,35,,34.27967690,131.52127420
852,"Yamanashi Prefecture",109,JP,Japan,19,,35.66351130,138.63888790
963,Ajloun,111,JO,Jordan,AJ,,32.33255840,35.75168440
965,Amman,111,JO,Jordan,AM,,31.94536330,35.92838950
959,Aqaba,111,JO,Jordan,AQ,,29.53208600,35.00628210
961,Balqa,111,JO,Jordan,BA,,32.03668060,35.72884800
960,Irbid,111,JO,Jordan,IR,,32.55696360,35.84789650
966,Jerash,111,JO,Jordan,JA,,32.27472370,35.89609540
956,Karak,111,JO,Jordan,KA,,31.18535270,35.70476820
964,Ma'an,111,JO,Jordan,MN,,30.19267890,35.72493190
958,Madaba,111,JO,Jordan,MD,,31.71960970,35.79327540
962,Mafraq,111,JO,Jordan,MA,,32.34169230,36.20201750
957,Tafilah,111,JO,Jordan,AT,,30.83380630,35.61605130
967,Zarqa,111,JO,Jordan,AZ,,32.06085050,36.09421210
145,"Akmola Region",112,KZ,Kazakhstan,AKM,,51.91653200,69.41104940
151,"Aktobe Region",112,KZ,Kazakhstan,AKT,,48.77970780,57.99743780
152,Almaty,112,KZ,Kazakhstan,ALA,,43.22201460,76.85124850
143,"Almaty Region",112,KZ,Kazakhstan,ALM,,45.01192270,78.42293920
153,"Atyrau Region",112,KZ,Kazakhstan,ATY,,47.10761880,51.91413300
155,Baikonur,112,KZ,Kazakhstan,BAY,,45.96458510,63.30524280
154,"East Kazakhstan Region",112,KZ,Kazakhstan,VOS,,48.70626870,80.79225340
147,"Jambyl Region",112,KZ,Kazakhstan,ZHA,,44.22203080,72.36579670
150,"Karaganda Region",112,KZ,Kazakhstan,KAR,,47.90221820,71.77068070
157,"Kostanay Region",112,KZ,Kazakhstan,KUS,,51.50770960,64.04790730
142,"Kyzylorda Region",112,KZ,Kazakhstan,KZY,,44.69226130,62.65718850
141,"Mangystau Region",112,KZ,Kazakhstan,MAN,,44.59080200,53.84995080
144,"North Kazakhstan Region",112,KZ,Kazakhstan,SEV,,54.16220660,69.93870710
156,Nur-Sultan,112,KZ,Kazakhstan,AST,,51.16052270,71.47035580
146,"Pavlodar Region",112,KZ,Kazakhstan,PAV,,52.28784440,76.97334530
149,"Turkestan Region",112,KZ,Kazakhstan,YUZ,,43.36669580,68.40944050
148,"West Kazakhstan Province",112,KZ,Kazakhstan,ZAP,,49.56797270,50.80666160
181,Baringo,113,KE,Kenya,01,county,0.85549880,36.08934060
210,Bomet,113,KE,Kenya,02,county,-0.80150090,35.30272260
168,Bungoma,113,KE,Kenya,03,county,0.56952520,34.55837660
161,Busia,113,KE,Kenya,04,county,0.43465060,34.24215970
201,Elgeyo-Marakwet,113,KE,Kenya,05,county,1.04982370,35.47819260
163,Embu,113,KE,Kenya,06,county,-0.65604770,37.72376780
196,Garissa,113,KE,Kenya,07,county,-0.45322930,39.64609880
195,"Homa Bay",113,KE,Kenya,08,county,-0.62206550,34.33103640
170,Isiolo,113,KE,Kenya,09,county,0.35243520,38.48499230
197,Kajiado,113,KE,Kenya,10,county,-2.09807510,36.78195050
158,Kakamega,113,KE,Kenya,11,county,0.30789400,34.77407930
193,Kericho,113,KE,Kenya,12,county,-0.18279130,35.47819260
199,Kiambu,113,KE,Kenya,13,county,-1.03137010,36.86807910
174,Kilifi,113,KE,Kenya,14,county,-3.51065080,39.90932690
167,Kirinyaga,113,KE,Kenya,15,county,-0.65905650,37.38272340
159,Kisii,113,KE,Kenya,16,county,-0.67733400,34.77960300
171,Kisumu,113,KE,Kenya,17,county,-0.09170160,34.76795680
211,Kitui,113,KE,Kenya,18,county,-1.68328220,38.31657250
173,Kwale,113,KE,Kenya,19,county,-4.18161150,39.46056120
164,Laikipia,113,KE,Kenya,20,county,0.36060630,36.78195050
166,Lamu,113,KE,Kenya,21,county,-2.23550580,40.47200040
184,Machakos,113,KE,Kenya,22,county,-1.51768370,37.26341460
188,Makueni,113,KE,Kenya,23,county,-2.25587340,37.89366630
187,Mandera,113,KE,Kenya,24,county,3.57379910,40.95868800
194,Marsabit,113,KE,Kenya,25,county,2.44264030,37.97845850
198,Meru,113,KE,Kenya,26,county,0.35571740,37.80876930
190,Migori,113,KE,Kenya,27,county,-0.93657020,34.41982430
200,Mombasa,113,KE,Kenya,28,county,-3.97682910,39.71371810
178,Murang'a,113,KE,Kenya,29,county,-0.78392810,37.04003390
191,"Nairobi City",113,KE,Kenya,30,county,-1.29206590,36.82194620
203,Nakuru,113,KE,Kenya,31,county,-0.30309880,36.08002600
165,Nandi,113,KE,Kenya,32,county,0.18358670,35.12687810
175,Narok,113,KE,Kenya,33,county,-1.10411100,36.08934060
209,Nyamira,113,KE,Kenya,34,county,-0.56694050,34.93412340
192,Nyandarua,113,KE,Kenya,35,county,-0.18038550,36.52296410
180,Nyeri,113,KE,Kenya,36,county,-0.41969150,37.04003390
207,Samburu,113,KE,Kenya,37,county,1.21545060,36.95410700
186,Siaya,113,KE,Kenya,38,county,-0.06173280,34.24215970
176,Taita–Taveta,113,KE,Kenya,39,county,-3.31606870,38.48499230
205,"Tana River",113,KE,Kenya,40,county,-1.65184680,39.65181650
185,Tharaka-Nithi,113,KE,Kenya,41,county,-0.29648510,37.72376780
183,"Trans Nzoia",113,KE,Kenya,42,county,1.05666670,34.95066250
206,Turkana,113,KE,Kenya,43,county,3.31224770,35.56578620
169,"Uasin Gishu",113,KE,Kenya,44,county,0.55276380,35.30272260
202,Vihiga,113,KE,Kenya,45,county,0.07675530,34.70776650
182,Wajir,113,KE,Kenya,46,county,1.63604750,40.30886260
208,"West Pokot",113,KE,Kenya,47,county,1.62100760,35.39050460
1831,"Gilbert Islands",114,KI,Kiribati,G,,0.35242620,174.75526340
1832,"Line Islands",114,KI,Kiribati,L,,1.74294390,-157.21328260
1830,"Phoenix Islands",114,KI,Kiribati,P,,33.32843690,-111.98247740
4876,"Đakovica District (Gjakove)",248,XK,Kosovo,XDG,,42.43757560,20.37854380
4877,"Gjilan District",248,XK,Kosovo,XGJ,,42.46352060,21.46940110
4878,"Kosovska Mitrovica District",248,XK,Kosovo,XKM,,42.89139090,20.86599950
3738,"Peć District",248,XK,Kosovo,XPE,,42.65921550,20.28876240
4879,"Pristina (Priştine)",248,XK,Kosovo,XPI,,42.66291380,21.16550280
3723,"Prizren District",248,XK,Kosovo,XPR,,42.21525220,20.74147720
4874,"Uroševac District (Ferizaj)",248,XK,Kosovo,XUF,,42.37018440,21.14832810
977,"Al Ahmadi",117,KW,Kuwait,AH,,28.57451250,48.10247430
975,"Al Farwaniyah",117,KW,Kuwait,FA,,29.27335700,47.94001540
972,"Al Jahra",117,KW,Kuwait,JA,,29.99318310,47.76347310
976,Capital,117,KW,Kuwait,KU,,26.22851610,50.58604970
973,Hawalli,117,KW,Kuwait,HA,,29.30567160,48.03076130
974,"Mubarak Al-Kabeer",117,KW,Kuwait,MU,,29.21224000,48.06051080
998,"Batken Region",118,KG,Kyrgyzstan,B,,39.97214250,69.85974060
1001,Bishkek,118,KG,Kyrgyzstan,GB,,42.87462120,74.56976170
1004,"Chuy Region",118,KG,Kyrgyzstan,C,,42.56550000,74.40566120
1002,"Issyk-Kul Region",118,KG,Kyrgyzstan,Y,,42.18594280,77.56194190
1000,"Jalal-Abad Region",118,KG,Kyrgyzstan,J,,41.10680800,72.89880690
999,"Naryn Region",118,KG,Kyrgyzstan,N,,41.29432270,75.34121790
1003,Osh,118,KG,Kyrgyzstan,GO,,36.06313990,-95.91828950
1005,"Osh Region",118,KG,Kyrgyzstan,O,,39.84073660,72.89880690
997,"Talas Region",118,KG,Kyrgyzstan,T,,42.28673390,72.52048270
982,"Attapeu Province",119,LA,Laos,AT,,14.93634000,107.10119310
991,"Bokeo Province",119,LA,Laos,BK,,20.28726620,100.70978670
985,"Bolikhamsai Province",119,LA,Laos,BL,,18.43629240,104.47233010
996,"Champasak Province",119,LA,Laos,CH,,14.65786640,105.96998780
989,"Houaphanh Province",119,LA,Laos,HO,,20.32541750,104.10013260
986,"Khammouane Province",119,LA,Laos,KH,,17.63840660,105.21948080
992,"Luang Namtha Province",119,LA,Laos,LM,,20.91701870,101.16173560
978,"Luang Prabang Province",119,LA,Laos,LP,,20.06562290,102.62162110
988,"Oudomxay Province",119,LA,Laos,OU,,20.49219290,101.88917210
987,"Phongsaly Province",119,LA,Laos,PH,,21.59193770,102.25479190
993,"Sainyabuli Province",119,LA,Laos,XA,,19.39078860,101.52480550
981,"Salavan Province",119,LA,Laos,SL,,15.81710730,106.25221430
990,"Savannakhet Province",119,LA,Laos,SV,,16.50653810,105.59433880
984,"Sekong Province",119,LA,Laos,XE,,15.57674460,107.00670310
979,"Vientiane Prefecture",119,LA,Laos,VT,,18.11054100,102.52980280
980,"Vientiane Province",119,LA,Laos,VI,,18.57050630,102.62162110
994,Xaisomboun,119,LA,Laos,XN,,,
983,"Xaisomboun Province",119,LA,Laos,XS,,18.43629240,104.47233010
995,"Xiangkhouang Province",119,LA,Laos,XI,,19.60930030,103.72891670
4445,"Aglona Municipality",120,LV,Latvia,001,,56.10890060,27.12862270
4472,"Aizkraukle Municipality",120,LV,Latvia,002,,56.64610800,25.23708540
4496,"Aizpute Municipality",120,LV,Latvia,003,,56.71825960,21.60727590
4499,"Aknīste Municipality",120,LV,Latvia,004,,56.16130370,25.74848270
4484,"Aloja Municipality",120,LV,Latvia,005,,57.76713600,24.87708390
4485,"Alsunga Municipality",120,LV,Latvia,006,,56.98285310,21.55559190
4487,"Alūksne Municipality",120,LV,Latvia,007,,57.42544850,27.04249680
4497,"Amata Municipality",120,LV,Latvia,008,,56.99387260,25.26276750
4457,"Ape Municipality",120,LV,Latvia,009,,57.53926970,26.69416490
4481,"Auce Municipality",120,LV,Latvia,010,,56.46016800,22.90547810
4427,"Babīte Municipality",120,LV,Latvia,012,,56.95415500,23.94539900
4482,"Baldone Municipality",120,LV,Latvia,013,,56.74246000,24.39115440
4498,"Baltinava Municipality",120,LV,Latvia,014,,56.94584680,27.64410660
4505,"Balvi Municipality",120,LV,Latvia,015,,57.13262400,27.26466850
4465,"Bauska Municipality",120,LV,Latvia,016,,56.41018680,24.20006890
4471,"Beverīna Municipality",120,LV,Latvia,017,,57.51971090,25.60736540
4468,"Brocēni Municipality",120,LV,Latvia,018,,56.73475410,22.63573710
4411,"Burtnieki Municipality",120,LV,Latvia,019,,57.69490040,25.27647770
4454,"Carnikava Municipality",120,LV,Latvia,020,,57.10241210,24.21086620
4469,"Cēsis Municipality",120,LV,Latvia,022,,57.31028970,25.26761250
4414,"Cesvaine Municipality",120,LV,Latvia,021,,56.96792640,26.30831720
4410,"Cibla Municipality",120,LV,Latvia,023,,56.61023440,27.86965980
4504,"Dagda Municipality",120,LV,Latvia,024,,56.09560890,27.53245900
4463,Daugavpils,120,LV,Latvia,DGV,,55.87473600,26.53617900
4492,"Daugavpils Municipality",120,LV,Latvia,025,,55.89917830,26.61020120
4437,"Dobele Municipality",120,LV,Latvia,026,,56.62630500,23.28090660
4428,"Dundaga Municipality",120,LV,Latvia,027,,57.50491670,22.35051140
4458,"Durbe Municipality",120,LV,Latvia,028,,56.62798570,21.49162450
4448,"Engure Municipality",120,LV,Latvia,029,,57.16235000,23.21966340
4444,"Ērgļi Municipality",120,LV,Latvia,030,,56.92370650,25.67538520
4510,"Garkalne Municipality",120,LV,Latvia,031,,57.01903870,24.38261810
4470,"Grobiņa Municipality",120,LV,Latvia,032,,56.53963200,21.16689200
4400,"Gulbene Municipality",120,LV,Latvia,033,,57.21556450,26.64529550
4441,"Iecava Municipality",120,LV,Latvia,034,,56.59867930,24.19962720
4511,"Ikšķile Municipality",120,LV,Latvia,035,,56.83736670,24.49747450
4399,"Ilūkste Municipality",120,LV,Latvia,036,,55.97825470,26.29650880
4449,"Inčukalns Municipality",120,LV,Latvia,037,,57.09943420,24.68555700
4475,"Jaunjelgava Municipality",120,LV,Latvia,038,,56.52836590,25.39214430
4407,"Jaunpiebalga Municipality",120,LV,Latvia,039,,57.14334710,25.99518880
4489,"Jaunpils Municipality",120,LV,Latvia,040,,56.73141940,23.01256160
4464,Jēkabpils,120,LV,Latvia,JKB,,56.50145500,25.87829900
4438,"Jēkabpils Municipality",120,LV,Latvia,042,,56.29193200,25.98120170
4500,Jelgava,120,LV,Latvia,JEL,,56.65110910,23.72135410
4424,"Jelgava Municipality",120,LV,Latvia,041,,56.58956890,23.66104810
4446,Jūrmala,120,LV,Latvia,JUR,,56.94707900,23.61684850
4420,"Kandava Municipality",120,LV,Latvia,043,,57.03406730,22.78018130
4453,"Kārsava Municipality",120,LV,Latvia,044,,56.76458420,27.73582950
4412,"Ķegums Municipality",120,LV,Latvia,051,,56.74753570,24.71736450
4435,"Ķekava Municipality",120,LV,Latvia,052,,56.80643510,24.19394930
4495,"Kocēni Municipality",120,LV,Latvia,045,,57.52262920,25.33495070
4452,"Koknese Municipality",120,LV,Latvia,046,,56.72055600,25.48939090
4474,"Krāslava Municipality",120,LV,Latvia,047,,55.89514640,27.18145770
4422,"Krimulda Municipality",120,LV,Latvia,048,,57.17912730,24.71401270
4413,"Krustpils Municipality",120,LV,Latvia,049,,56.54155780,26.24463970
4490,"Kuldīga Municipality",120,LV,Latvia,050,,56.96872570,21.96137390
4512,"Lielvārde Municipality",120,LV,Latvia,053,,56.73929760,24.97116180
4460,Liepāja,120,LV,Latvia,LPX,,56.50466780,21.01080600
4488,"Līgatne Municipality",120,LV,Latvia,055,,57.19442040,25.09406810
4418,"Limbaži Municipality",120,LV,Latvia,054,,57.54032270,24.71344510
4401,"Līvāni Municipality",120,LV,Latvia,056,,56.35509420,26.17251900
4419,"Lubāna Municipality",120,LV,Latvia,057,,56.89992690,26.71987890
4501,"Ludza Municipality",120,LV,Latvia,058,,56.54595900,27.71431990
4433,"Madona Municipality",120,LV,Latvia,059,,56.85989230,26.22762010
4461,"Mālpils Municipality",120,LV,Latvia,061,,57.00841190,24.95742780
4450,"Mārupe Municipality",120,LV,Latvia,062,,56.89657170,24.04600490
4513,"Mazsalaca Municipality",120,LV,Latvia,060,,57.92677490,25.06698950
4451,"Mērsrags Municipality",120,LV,Latvia,063,,57.33068810,23.10237070
4398,"Naukšēni Municipality",120,LV,Latvia,064,,57.92953610,25.51192660
4432,"Nereta Municipality",120,LV,Latvia,065,,56.19866550,25.32529690
4436,"Nīca Municipality",120,LV,Latvia,066,,56.34649830,21.06549300
4416,"Ogre Municipality",120,LV,Latvia,067,,56.81473550,24.60445550
4417,"Olaine Municipality",120,LV,Latvia,068,,56.79523530,24.01535890
4442,"Ozolnieki Municipality",120,LV,Latvia,069,,56.67563050,23.89948160
4507,"Pārgauja Municipality",120,LV,Latvia,070,,57.36481220,24.98220450
4467,"Pāvilosta Municipality",120,LV,Latvia,071,,56.88654240,21.19468490
4405,"Pļaviņas Municipality",120,LV,Latvia,072,,56.61773130,25.71940430
4483,"Preiļi Municipality",120,LV,Latvia,073,,56.15111570,26.74397670
4429,"Priekule Municipality",120,LV,Latvia,074,,56.41794130,21.55033360
4506,"Priekuļi Municipality",120,LV,Latvia,075,,57.36171380,25.44104230
4479,"Rauna Municipality",120,LV,Latvia,076,,57.33169300,25.61003390
4509,Rēzekne,120,LV,Latvia,REZ,,56.50992230,27.33313570
4455,"Rēzekne Municipality",120,LV,Latvia,077,,56.32736380,27.32843310
4502,"Riebiņi Municipality",120,LV,Latvia,078,,56.34361900,26.80181380
4491,Riga,120,LV,Latvia,RIX,,56.94964870,24.10518650
4440,"Roja Municipality",120,LV,Latvia,079,,57.50467130,22.80121640
4493,"Ropaži Municipality",120,LV,Latvia,080,,56.96157860,24.60104760
4503,"Rucava Municipality",120,LV,Latvia,081,,56.15931240,21.16181210
4423,"Rugāji Municipality",120,LV,Latvia,082,,57.00560230,27.13172030
4426,"Rūjiena Municipality",120,LV,Latvia,084,,57.89372910,25.33910080
4404,"Rundāle Municipality",120,LV,Latvia,083,,56.40972100,24.01241390
4434,"Sala Municipality",120,LV,Latvia,085,,59.96796130,16.49782170
4396,"Salacgrīva Municipality",120,LV,Latvia,086,,57.75808830,24.35431810
4402,"Salaspils Municipality",120,LV,Latvia,087,,56.86081520,24.34978810
4439,"Saldus Municipality",120,LV,Latvia,088,,56.66650880,22.49354930
4443,"Saulkrasti Municipality",120,LV,Latvia,089,,57.25794180,24.41831460
4408,"Sēja Municipality",120,LV,Latvia,090,,57.20069950,24.59228210
4476,"Sigulda Municipality",120,LV,Latvia,091,,57.10550920,24.83142590
4415,"Skrīveri Municipality",120,LV,Latvia,092,,56.67613910,25.09788490
4447,"Skrunda Municipality",120,LV,Latvia,093,,56.66434580,22.00457290
4462,"Smiltene Municipality",120,LV,Latvia,094,,57.42303320,25.90027800
4478,"Stopiņi Municipality",120,LV,Latvia,095,,56.93644900,24.28729490
4494,"Strenči Municipality",120,LV,Latvia,096,,57.62254710,25.80480860
4459,"Talsi Municipality",120,LV,Latvia,097,,57.34152080,22.57131250
4480,"Tērvete Municipality",120,LV,Latvia,098,,56.41192010,23.31883320
4409,"Tukums Municipality",120,LV,Latvia,099,,56.96728680,23.15243790
4508,"Vaiņode Municipality",120,LV,Latvia,100,,56.41542710,21.85139840
4425,"Valka Municipality",120,LV,Latvia,101,,57.77439000,26.01700500
4473,Valmiera,120,LV,Latvia,VMR,,57.53846590,25.42636180
4431,"Varakļāni Municipality",120,LV,Latvia,102,,56.66880060,26.56364140
4406,"Vārkava Municipality",120,LV,Latvia,103,,56.24657440,26.56643710
4466,"Vecpiebalga Municipality",120,LV,Latvia,104,,57.06033560,25.81615920
4397,"Vecumnieki Municipality",120,LV,Latvia,105,,56.60623370,24.52218910
4421,Ventspils,120,LV,Latvia,VEN,,57.39372160,21.56470660
4403,"Ventspils Municipality",120,LV,Latvia,106,,57.28336820,21.85875580
4456,"Viesīte Municipality",120,LV,Latvia,107,,56.31130850,25.50704640
4477,"Viļaka Municipality",120,LV,Latvia,108,,57.17222630,27.66731880
4486,"Viļāni Municipality",120,LV,Latvia,109,,56.54561710,26.91679270
4430,"Zilupe Municipality",120,LV,Latvia,110,,56.30189850,28.13395900
2285,Akkar,121,LB,Lebanon,AK,,34.53287630,36.13281320
2283,Baalbek-Hermel,121,LB,Lebanon,BH,,34.26585560,36.34980970
2286,Beirut,121,LB,Lebanon,BA,,33.88861060,35.49547720
2287,Beqaa,121,LB,Lebanon,BI,,33.84626620,35.90194890
2282,"Mount Lebanon",121,LB,Lebanon,JL,,33.81008580,35.59731390
2288,Nabatieh,121,LB,Lebanon,NA,,33.37716930,35.48382930
2284,North,121,LB,Lebanon,AS,,34.43806250,35.83082330
2281,South,121,LB,Lebanon,JA,,33.27214790,35.20327780
3030,"Berea District",122,LS,Lesotho,D,,41.36616140,-81.85430260
3029,"Butha-Buthe District",122,LS,Lesotho,B,,-28.76537540,28.24681480
3026,"Leribe District",122,LS,Lesotho,C,,-28.86380650,28.04788260
3022,"Mafeteng District",122,LS,Lesotho,E,,-29.80410080,27.50261740
3028,"Maseru District",122,LS,Lesotho,A,,-29.51656500,27.83114280
3023,"Mohale's Hoek District",122,LS,Lesotho,F,,-30.14259170,27.46738450
3024,"Mokhotlong District",122,LS,Lesotho,J,,-29.25731930,28.95286450
3025,"Qacha's Nek District",122,LS,Lesotho,H,,-30.11145650,28.67897900
3027,"Quthing District",122,LS,Lesotho,G,,-30.40156870,27.70801330
3031,"Thaba-Tseka District",122,LS,Lesotho,K,,-29.52389750,28.60897520
3041,"Bomi County",123,LR,Liberia,BM,,6.75629260,-10.84514670
3034,"Bong County",123,LR,Liberia,BG,,6.82950190,-9.36730840
3044,"Gbarpolu County",123,LR,Liberia,GP,,7.49526370,-10.08072980
3040,"Grand Bassa County",123,LR,Liberia,GB,,6.23084520,-9.81249350
3036,"Grand Cape Mount County",123,LR,Liberia,CM,,7.04677580,-11.07117580
3039,"Grand Gedeh County",123,LR,Liberia,GG,,5.92220780,-8.22129790
3045,"Grand Kru County",123,LR,Liberia,GK,,4.76138620,-8.22129790
3037,"Lofa County",123,LR,Liberia,LO,,8.19111840,-9.72326730
3043,"Margibi County",123,LR,Liberia,MG,,6.51518750,-10.30488970
3042,"Maryland County",123,LR,Liberia,MY,,39.04575490,-76.64127120
3032,"Montserrado County",123,LR,Liberia,MO,,6.55258150,-10.52961150
3046,Nimba,123,LR,Liberia,NI,,7.61666670,-8.41666670
3033,"River Cess County",123,LR,Liberia,RI,,5.90253280,-9.45615500
3038,"River Gee County",123,LR,Liberia,RG,,5.26048940,-7.87216000
3035,"Sinoe County",123,LR,Liberia,SI,,5.49871000,-8.66005860
2964,"Al Wahat District",124,LY,Libya,WA,,29.04668080,21.85685860
2981,Benghazi,124,LY,Libya,BA,,32.11942420,20.08679090
2966,"Derna District",124,LY,Libya,DR,,32.75561300,22.63774320
2969,"Ghat District",124,LY,Libya,GT,,24.96403710,10.17592850
2980,"Jabal al Akhdar",124,LY,Libya,JA,,23.18560810,57.37138790
2974,"Jabal al Gharbi District",124,LY,Libya,JG,,30.26380320,12.80547530
2979,Jafara,124,LY,Libya,JI,,32.45259040,12.94355360
2970,Jufra,124,LY,Libya,JU,,27.98351350,16.91225100
2972,"Kufra District",124,LY,Libya,KF,,23.31123890,21.85685860
2968,"Marj District",124,LY,Libya,MJ,,32.05503630,21.18911510
2978,"Misrata District",124,LY,Libya,MI,,32.32558840,15.09925560
2961,Murqub,124,LY,Libya,MB,,32.45996770,14.10013260
2967,"Murzuq District",124,LY,Libya,MQ,,25.91822620,13.92600010
2976,"Nalut District",124,LY,Libya,NL,,31.87423480,10.97504840
2962,"Nuqat al Khams",124,LY,Libya,NQ,,32.69149090,11.88917210
2965,"Sabha District",124,LY,Libya,SB,,27.03654060,14.42902360
2977,"Sirte District",124,LY,Libya,SR,,31.18968900,16.57019270
2971,"Tripoli District",124,LY,Libya,TB,,32.64080210,13.26634790
2973,"Wadi al Hayaa District",124,LY,Libya,WD,,26.42259260,12.62162110
2975,"Wadi al Shatii District",124,LY,Libya,WS,,27.73514680,12.43805810
2963,"Zawiya District",124,LY,Libya,ZA,,32.76302820,12.73649620
458,Balzers,125,LI,Liechtenstein,01,,42.05283570,-88.03668480
451,Eschen,125,LI,Liechtenstein,02,,40.76695740,-73.95228210
457,Gamprin,125,LI,Liechtenstein,03,,47.21324900,9.50251950
455,Mauren,125,LI,Liechtenstein,04,,47.21892850,9.54173500
454,Planken,125,LI,Liechtenstein,05,,40.66505760,-73.50479800
453,Ruggell,125,LI,Liechtenstein,06,,47.25292220,9.54021270
450,Schaan,125,LI,Liechtenstein,07,,47.12043400,9.59416020
449,Schellenberg,125,LI,Liechtenstein,08,,47.23096600,9.54678430
459,Triesen,125,LI,Liechtenstein,09,,47.10979880,9.52482960
456,Triesenberg,125,LI,Liechtenstein,10,,47.12245110,9.57019850
452,Vaduz,125,LI,Liechtenstein,11,,47.14103030,9.52092770
1561,"Akmenė District Municipality",126,LT,Lithuania,01,,56.24550290,22.74711690
1605,"Alytus City Municipality",126,LT,Lithuania,02,,54.39629380,24.04587610
1574,"Alytus County",126,LT,Lithuania,AL,,54.20002140,24.15126340
1599,"Alytus District Municipality",126,LT,Lithuania,03,,54.32974960,24.19609310
1603,"Birštonas Municipality",126,LT,Lithuania,05,,54.56696640,24.00930980
1566,"Biržai District Municipality",126,LT,Lithuania,06,,56.20177190,24.75601180
1579,"Druskininkai municipality",126,LT,Lithuania,07,,53.99336850,24.03424380
1559,"Elektrėnai municipality",126,LT,Lithuania,08,,54.76539340,24.77405830
1562,"Ignalina District Municipality",126,LT,Lithuania,09,,55.40903420,26.32848930
1567,"Jonava District Municipality",126,LT,Lithuania,10,,55.07272420,24.27933370
1581,"Joniškis District Municipality",126,LT,Lithuania,11,,56.23607300,23.61365790
1555,"Jurbarkas District Municipality",126,LT,Lithuania,12,,55.07740700,22.74195690
1583,"Kaišiadorys District Municipality",126,LT,Lithuania,13,,54.85886690,24.42779290
1591,"Kalvarija municipality",126,LT,Lithuania,14,,54.37616740,23.19203210
1580,"Kaunas City Municipality",126,LT,Lithuania,15,,54.91453260,23.90535180
1556,"Kaunas County",126,LT,Lithuania,KU,,54.98728630,23.95257360
1565,"Kaunas District Municipality",126,LT,Lithuania,16,,54.99362360,23.63249410
1575,"Kazlų Rūda municipality",126,LT,Lithuania,17,,54.78075260,23.48402430
1584,"Kėdainiai District Municipality",126,LT,Lithuania,18,,55.35609470,23.98326830
1618,"Kelmė District Municipality",126,LT,Lithuania,19,,55.62663520,22.87817200
1597,"Klaipeda City Municipality",126,LT,Lithuania,20,,55.70329480,21.14427950
1600,"Klaipėda County",126,LT,Lithuania,KL,,55.65197440,21.37439560
1604,"Klaipėda District Municipality",126,LT,Lithuania,21,,55.68416150,21.44164640
1571,"Kretinga District Municipality",126,LT,Lithuania,22,,55.88384200,21.23509190
1585,"Kupiškis District Municipality",126,LT,Lithuania,23,,55.84287410,25.02958160
1611,"Lazdijai District Municipality",126,LT,Lithuania,24,,54.23432670,23.51565050
1570,"Marijampolė County",126,LT,Lithuania,MR,,54.78199710,23.13413650
1610,"Marijampolė Municipality",126,LT,Lithuania,25,,54.57110940,23.48593710
1557,"Mažeikiai District Municipality",126,LT,Lithuania,26,,56.30924390,22.34146800
1582,"Molėtai District Municipality",126,LT,Lithuania,27,,55.22653090,25.41800110
1563,"Neringa Municipality",126,LT,Lithuania,28,,55.45724030,21.08390050
1612,"Pagėgiai municipality",126,LT,Lithuania,29,,55.17213200,21.96836140
1595,"Pakruojis District Municipality",126,LT,Lithuania,30,,56.07326050,23.93899060
1588,"Palanga City Municipality",126,LT,Lithuania,31,,55.92019800,21.06776140
1589,"Panevėžys City Municipality",126,LT,Lithuania,32,,55.73479150,24.35747740
1558,"Panevėžys County",126,LT,Lithuania,PN,,55.97480490,25.07947670
1614,"Panevėžys District Municipality",126,LT,Lithuania,33,,55.61667280,24.31422830
1616,"Pasvalys District Municipality",126,LT,Lithuania,34,,56.06046190,24.39629100
1553,"Plungė District Municipality",126,LT,Lithuania,35,,55.91078400,21.84540690
1578,"Prienai District Municipality",126,LT,Lithuania,36,,54.63835800,23.94680090
1568,"Radviliškis District Municipality",126,LT,Lithuania,37,,55.81083990,23.54648700
1587,"Raseiniai District Municipality",126,LT,Lithuania,38,,55.38194990,23.11561290
1590,"Rietavas municipality",126,LT,Lithuania,39,,55.70217190,21.99865640
1615,"Rokiškis District Municipality",126,LT,Lithuania,40,,55.95550390,25.58592490
1576,"Šakiai District Municipality",126,LT,Lithuania,41,,54.95267100,23.04801990
1577,"Šalčininkai District Municipality",126,LT,Lithuania,42,,54.30976700,25.38756400
1609,"Šiauliai City Municipality",126,LT,Lithuania,43,,55.93490850,23.31368230
1586,"Šiauliai County",126,LT,Lithuania,SA,,55.99857510,23.13800510
1554,"Šiauliai District Municipality",126,LT,Lithuania,44,,55.97214560,23.03323710
1613,"Šilalė District Municipality",126,LT,Lithuania,45,,55.49268000,22.18455590
1607,"Šilutė District Municipality",126,LT,Lithuania,46,,55.35041400,21.46598590
1594,"Širvintos District Municipality",126,LT,Lithuania,47,,55.04310200,24.95698100
1617,"Skuodas District Municipality",126,LT,Lithuania,48,,56.27021690,21.52143310
1560,"Švenčionys District Municipality",126,LT,Lithuania,49,,55.10280980,26.00718550
1573,"Tauragė County",126,LT,Lithuania,TA,,55.30725860,22.35729390
1572,"Tauragė District Municipality",126,LT,Lithuania,50,,55.25036600,22.29095000
1569,"Telšiai County",126,LT,Lithuania,TE,,56.10266160,22.11139150
1608,"Telšiai District Municipality",126,LT,Lithuania,51,,55.91752150,22.34518400
1593,"Trakai District Municipality",126,LT,Lithuania,52,,54.63791130,24.93468940
1596,"Ukmergė District Municipality",126,LT,Lithuania,53,,55.24526500,24.77607490
1621,"Utena County",126,LT,Lithuania,UT,,55.53189690,25.79046990
1598,"Utena District Municipality",126,LT,Lithuania,54,,55.50846140,25.68326420
1602,"Varėna District Municipality",126,LT,Lithuania,55,,54.22033300,24.57899700
1620,"Vilkaviškis District Municipality",126,LT,Lithuania,56,,54.65194500,23.03515500
1606,"Vilnius City Municipality",126,LT,Lithuania,57,,54.67107610,25.28787210
1601,"Vilnius County",126,LT,Lithuania,VL,,54.80865020,25.21821390
1592,"Vilnius District Municipality",126,LT,Lithuania,58,,54.77325780,25.58671130
1564,"Visaginas Municipality",126,LT,Lithuania,59,,55.59411800,26.43079540
1619,"Zarasai District Municipality",126,LT,Lithuania,60,,55.73096090,26.24529500
1518,"Canton of Capellen",127,LU,Luxembourg,CA,,49.64039310,5.95538460
1521,"Canton of Clervaux",127,LU,Luxembourg,CL,,50.05463130,6.02858750
1513,"Canton of Diekirch",127,LU,Luxembourg,DI,,49.86717840,6.15956330
1515,"Canton of Echternach",127,LU,Luxembourg,EC,,49.81141330,6.41756350
1517,"Canton of Esch-sur-Alzette",127,LU,Luxembourg,ES,,49.50088050,5.98609250
1525,"Canton of Grevenmacher",127,LU,Luxembourg,GR,,49.68084100,6.44075930
1527,"Canton of Luxembourg",127,LU,Luxembourg,LU,,49.63010250,6.15201850
1522,"Canton of Mersch",127,LU,Luxembourg,ME,,49.75429060,6.12921850
1516,"Canton of Redange",127,LU,Luxembourg,RD,,49.76455000,5.88948000
1519,"Canton of Remich",127,LU,Luxembourg,RM,,49.54501700,6.36742220
1523,"Canton of Vianden",127,LU,Luxembourg,VD,,49.93419240,6.20199170
1526,"Canton of Wiltz",127,LU,Luxembourg,WI,,49.96622000,5.93243060
1524,"Diekirch District",127,LU,Luxembourg,D,,49.86717200,6.15963620
1520,"Grevenmacher District",127,LU,Luxembourg,G,,49.68085100,6.44075240
1514,"Luxembourg District",127,LU,Luxembourg,L,,49.59537060,6.13331780
2951,"Antananarivo Province",130,MG,Madagascar,T,,-18.70514740,46.82528380
2950,"Antsiranana Province",130,MG,Madagascar,D,,-13.77153900,49.52799960
2948,"Fianarantsoa Province",130,MG,Madagascar,F,,-22.35362400,46.82528380
2953,"Mahajanga Province",130,MG,Madagascar,M,,-16.52388300,46.51626200
2952,"Toamasina Province",130,MG,Madagascar,A,,-18.14428110,49.39578360
2949,"Toliara Province",130,MG,Madagascar,U,,-23.35161910,43.68549360
3096,"Balaka District",131,MW,Malawi,BA,,-15.05065950,35.08285880
3102,"Blantyre District",131,MW,Malawi,BL,,-15.67785410,34.95066250
3092,"Central Region",131,MW,Malawi,C,,,
3107,"Chikwawa District",131,MW,Malawi,CK,,-16.19584460,34.77407930
3109,"Chiradzulu District",131,MW,Malawi,CR,,-15.74231510,35.25879640
3087,"Chitipa district",131,MW,Malawi,CT,,-9.70376550,33.27002530
3097,"Dedza District",131,MW,Malawi,DE,,-14.18945850,34.24215970
3090,"Dowa District",131,MW,Malawi,DO,,-13.60412560,33.88577470
3091,"Karonga District",131,MW,Malawi,KR,,-9.90363650,33.97500180
3094,"Kasungu District",131,MW,Malawi,KS,,-13.13670650,33.25879300
3093,"Likoma District",131,MW,Malawi,LK,,-12.05840050,34.73540310
3101,"Lilongwe District",131,MW,Malawi,LI,,-14.04752280,33.61757700
3082,"Machinga District",131,MW,Malawi,MH,,-14.94072630,35.47819260
3110,"Mangochi District",131,MW,Malawi,MG,,-14.13882480,35.03881640
3099,"Mchinji District",131,MW,Malawi,MC,,-13.74015250,32.98883190
3103,"Mulanje District",131,MW,Malawi,MU,,-15.93464340,35.52200120
3084,"Mwanza District",131,MW,Malawi,MW,,-2.46711970,32.89868120
3104,"Mzimba District",131,MW,Malawi,MZ,,-11.74754520,33.52800720
3095,"Nkhata Bay District",131,MW,Malawi,NB,,-11.71850420,34.33103640
3100,"Nkhotakota District",131,MW,Malawi,NK,,-12.75419610,34.24215970
3105,"Northern Region",131,MW,Malawi,N,,,
3085,"Nsanje District",131,MW,Malawi,NS,,-16.72882020,35.17087410
3088,"Ntcheu District",131,MW,Malawi,NU,,-14.90375380,34.77407930
3111,"Ntchisi District",131,MW,Malawi,NI,,-13.28419920,33.88577470
3108,"Phalombe District",131,MW,Malawi,PH,,-15.70920380,35.65328480
3089,"Rumphi District",131,MW,Malawi,RU,,-10.78515370,34.33103640
3086,"Salima District",131,MW,Malawi,SA,,-13.68095860,34.41982430
3106,"Southern Region",131,MW,Malawi,S,,32.75049570,-97.33154760
3098,"Thyolo District",131,MW,Malawi,TH,,-16.12991770,35.12687810
3083,"Zomba District",131,MW,Malawi,ZO,,-15.37658570,35.33565180
1950,Johor,132,MY,Malaysia,01,,1.48536820,103.76181540
1947,Kedah,132,MY,Malaysia,02,,6.11839640,100.36845950
1946,Kelantan,132,MY,Malaysia,03,,6.12539690,102.23807100
1949,"Kuala Lumpur",132,MY,Malaysia,14,,3.13900300,101.68685500
1935,Labuan,132,MY,Malaysia,15,,5.28314560,115.23082500
1941,Malacca,132,MY,Malaysia,04,,2.18959400,102.25008680
1948,"Negeri Sembilan",132,MY,Malaysia,05,,2.72580580,101.94237820
1940,Pahang,132,MY,Malaysia,06,,3.81263180,103.32562040
1939,Penang,132,MY,Malaysia,07,,5.41639350,100.33267860
1943,Perak,132,MY,Malaysia,08,,4.59211260,101.09010900
1938,Perlis,132,MY,Malaysia,09,,29.92270940,-90.12285590
1945,Putrajaya,132,MY,Malaysia,16,,2.92636100,101.69644500
1936,Sabah,132,MY,Malaysia,12,,5.97883980,116.07531990
1937,Sarawak,132,MY,Malaysia,13,,1.55327830,110.35921270
1944,Selangor,132,MY,Malaysia,10,,3.07383790,101.51834690
1942,Terengganu,132,MY,Malaysia,11,,5.31169160,103.13241540
2594,"Addu Atoll",133,MV,Maldives,01,,-0.63009950,73.15856260
2587,"Alif Alif Atoll",133,MV,Maldives,02,,4.08500000,72.85154790
2600,"Alif Dhaal Atoll",133,MV,Maldives,00,,3.65433020,72.80427970
2604,"Central Province",133,MV,Maldives,CE,,,
2590,"Dhaalu Atoll",133,MV,Maldives,17,,2.84685020,72.94605660
2599,"Faafu Atoll",133,MV,Maldives,14,,3.23094090,72.94605660
2598,"Gaafu Alif Atoll",133,MV,Maldives,27,,0.61248130,73.32370800
2603,"Gaafu Dhaalu Atoll",133,MV,Maldives,28,,0.35880400,73.18216230
2595,"Gnaviyani Atoll",133,MV,Maldives,29,,-0.30064250,73.42391430
2586,"Haa Alif Atoll",133,MV,Maldives,07,,6.99034880,72.94605660
2597,"Haa Dhaalu Atoll",133,MV,Maldives,23,,6.57827170,72.94605660
2596,"Kaafu Atoll",133,MV,Maldives,26,,4.45589790,73.55941280
2601,"Laamu Atoll",133,MV,Maldives,05,,1.94307370,73.41802110
2607,"Lhaviyani Atoll",133,MV,Maldives,03,,5.37470210,73.51229280
2609,Malé,133,MV,Maldives,MLE,,46.34888670,10.90724890
2608,"Meemu Atoll",133,MV,Maldives,12,,3.00903450,73.51229280
2592,"Noonu Atoll",133,MV,Maldives,25,,5.85512760,73.32370800
2589,"North Central Province",133,MV,Maldives,NC,,,
2588,"North Province",133,MV,Maldives,NO,,8.88550270,80.27673270
2602,"Raa Atoll",133,MV,Maldives,13,,5.60064570,72.94605660
2585,"Shaviyani Atoll",133,MV,Maldives,24,,6.17511000,73.13496050
2606,"South Central Province",133,MV,Maldives,SC,,7.25649960,80.72144170
2605,"South Province",133,MV,Maldives,SU,,-21.74820060,166.17837390
2591,"Thaa Atoll",133,MV,Maldives,08,,2.43111610,73.18216230
2593,"Upper South Province",133,MV,Maldives,US,,0.23070000,73.27948460
2584,"Vaavu Atoll",133,MV,Maldives,04,,3.39554380,73.51229280
253,Bamako,134,ML,Mali,BKO,,12.63923160,-8.00288920
258,"Gao Region",134,ML,Mali,7,,16.90663320,1.52086240
252,"Kayes Region",134,ML,Mali,1,,14.08183080,-9.90181310
257,"Kidal Region",134,ML,Mali,8,,18.79868320,1.83183340
250,"Koulikoro Region",134,ML,Mali,2,,13.80180740,-7.43813550
251,"Ménaka Region",134,ML,Mali,9,,15.91564210,2.39617400
255,"Mopti Region",134,ML,Mali,5,,14.63380390,-3.41955270
249,"Ségou Region",134,ML,Mali,4,,13.83944560,-6.06791940
254,"Sikasso Region",134,ML,Mali,3,,10.89051860,-7.43813550
256,"Taoudénit Region",134,ML,Mali,10,,22.67641320,-3.97891430
248,"Tombouctou Region",134,ML,Mali,6,,21.05267060,-3.74350900
110,Attard,135,MT,Malta,01,,35.89049670,14.41993220
108,Balzan,135,MT,Malta,02,,35.89574140,14.45340650
107,Birgu,135,MT,Malta,03,,35.88792140,14.52256200
97,Birkirkara,135,MT,Malta,04,,35.89547060,14.46650720
88,Birżebbuġa,135,MT,Malta,05,,35.81359890,14.52474630
138,Cospicua,135,MT,Malta,06,,35.88067530,14.52183380
117,Dingli,135,MT,Malta,07,,35.86273090,14.38501070
129,Fgura,135,MT,Malta,08,,35.87382690,14.52329010
84,Floriana,135,MT,Malta,09,,45.49521850,-73.71395760
134,Fontana,135,MT,Malta,10,,34.09223350,-117.43504800
130,Għajnsielem,135,MT,Malta,13,,36.02479660,14.28029610
92,Għarb,135,MT,Malta,14,,36.06890900,14.20180980
120,Għargħur,135,MT,Malta,15,,35.92205690,14.45631760
106,Għasri,135,MT,Malta,16,,36.06680750,14.21924750
124,Għaxaq,135,MT,Malta,17,,35.84403590,14.51600900
118,Gudja,135,MT,Malta,11,,35.84698030,14.50290400
113,Gżira,135,MT,Malta,12,,35.90589700,14.49533380
105,Ħamrun,135,MT,Malta,18,,35.88612370,14.48834420
93,Iklin,135,MT,Malta,19,,35.90987740,14.45777320
99,Kalkara,135,MT,Malta,21,,35.89142420,14.53202780
91,Kerċem,135,MT,Malta,22,,36.04479390,14.22506050
82,Kirkop,135,MT,Malta,23,,35.84378620,14.48543240
126,Lija,135,MT,Malta,24,,49.18007600,-123.10331700
77,Luqa,135,MT,Malta,25,,35.85828650,14.48688830
128,Marsa,135,MT,Malta,26,,34.03195870,-118.44555350
137,Marsaskala,135,MT,Malta,27,,35.86036400,14.55678760
78,Marsaxlokk,135,MT,Malta,28,,35.84116990,14.53930970
89,Mdina,135,MT,Malta,29,,35.88809300,14.40683570
102,Mellieħa,135,MT,Malta,30,,35.95235290,14.35009750
109,Mġarr,135,MT,Malta,31,,35.91893270,14.36173430
140,Mosta,135,MT,Malta,32,,35.91415040,14.42284270
74,Mqabba,135,MT,Malta,33,,35.84441430,14.46941860
96,Msida,135,MT,Malta,34,,35.89563880,14.48688830
131,Mtarfa,135,MT,Malta,35,,35.88951250,14.39519530
132,Munxar,135,MT,Malta,36,,36.02880580,14.22506050
133,Nadur,135,MT,Malta,37,,36.04470190,14.29192730
112,Naxxar,135,MT,Malta,38,,35.93175180,14.43157460
115,Paola,135,MT,Malta,39,,38.57223530,-94.87912940
125,Pembroke,135,MT,Malta,40,,34.68016260,-79.19503730
127,Pietà,135,MT,Malta,41,,42.21862000,-83.73464700
79,Qala,135,MT,Malta,42,,36.03886280,14.31810100
119,Qormi,135,MT,Malta,43,,35.87643880,14.46941860
111,Qrendi,135,MT,Malta,44,,35.83284880,14.45486210
83,Rabat,135,MT,Malta,46,,33.97159040,-6.84981290
87,"Saint Lawrence",135,MT,Malta,50,,38.95780560,-95.25656890
75,"San Ġwann",135,MT,Malta,49,,35.90773650,14.47524160
116,Sannat,135,MT,Malta,52,,36.01926430,14.25994370
94,"Santa Luċija",135,MT,Malta,53,,35.85614200,14.50436000
90,"Santa Venera",135,MT,Malta,54,,35.89022010,14.47669740
136,Senglea,135,MT,Malta,20,,35.88730410,14.51673710
98,Siġġiewi,135,MT,Malta,55,,35.84637420,14.43157460
104,Sliema,135,MT,Malta,56,,35.91100810,14.50290400
100,"St. Julian's",135,MT,Malta,48,,42.21225130,-85.89171270
139,"St. Paul's Bay",135,MT,Malta,51,,35.93601700,14.39665030
86,Swieqi,135,MT,Malta,57,,35.91911820,14.46941860
122,"Ta' Xbiex",135,MT,Malta,58,,35.89914480,14.49635190
103,Tarxien,135,MT,Malta,59,,35.86725520,14.51164050
95,Valletta,135,MT,Malta,60,,35.89890850,14.51455280
101,Victoria,135,MT,Malta,45,,28.80526740,-97.00359820
114,Xagħra,135,MT,Malta,61,,36.05084500,14.26748200
121,Xewkija,135,MT,Malta,62,,36.02992360,14.25994370
81,Xgħajra,135,MT,Malta,63,,35.88682820,14.54723910
123,Żabbar,135,MT,Malta,64,,35.87247150,14.54513540
85,"Żebbuġ Gozo",135,MT,Malta,65,,36.07164030,14.24540800
80,"Żebbuġ Malta",135,MT,Malta,66,,35.87646480,14.43908400
135,Żejtun,135,MT,Malta,67,,35.85487140,14.53639690
76,Żurrieq,135,MT,Malta,68,,35.82163060,14.48106480
2574,"Ralik Chain",137,MH,"Marshall Islands",L,,8.13614600,164.88679560
2573,"Ratak Chain",137,MH,"Marshall Islands",T,,10.27632760,170.55009370
3344,Adrar,139,MR,Mauritania,07,region,19.86521760,-12.80547530
3349,Assaba,139,MR,Mauritania,03,region,16.77595580,-11.52480550
3339,Brakna,139,MR,Mauritania,05,region,17.23175610,-13.17403480
3346,"Dakhlet Nouadhibou",139,MR,Mauritania,08,region,20.59855880,-16.25221430
3341,Gorgol,139,MR,Mauritania,04,region,15.97173570,-12.62162110
3350,Guidimaka,139,MR,Mauritania,10,region,15.25573310,-12.25479190
3338,"Hodh Ech Chargui",139,MR,Mauritania,01,region,18.67370260,-7.09287700
3351,"Hodh El Gharbi",139,MR,Mauritania,02,region,16.69121490,-9.54509740
3342,Inchiri,139,MR,Mauritania,12,region,20.02805610,-15.40680790
3343,Nouakchott-Nord,139,MR,Mauritania,14,region,18.11302050,-15.89949560
3352,Nouakchott-Ouest,139,MR,Mauritania,13,region,18.15113570,-15.99349100
3347,Nouakchott-Sud,139,MR,Mauritania,15,region,17.97092880,-15.94648740
3345,Tagant,139,MR,Mauritania,09,region,18.54675270,-9.90181310
3340,"Tiris Zemmour",139,MR,Mauritania,11,region,24.57737640,-9.90181310
3348,Trarza,139,MR,Mauritania,06,region,17.86649640,-14.65878210
3248,"Agalega Islands",140,MU,Mauritius,AG,dependency,-10.40000000,56.61666670
3259,"Black River",140,MU,Mauritius,BL,district,-20.37084920,57.39486490
3254,Flacq,140,MU,Mauritius,FL,district,-20.22578360,57.71192740
3264,"Grand Port",140,MU,Mauritius,GP,district,-20.38515460,57.66657420
3253,Moka,140,MU,Mauritius,MO,district,-20.23997820,57.57592600
3250,Pamplemousses,140,MU,Mauritius,PA,district,-20.11360080,57.57592600
3263,"Plaines Wilhems",140,MU,Mauritius,PW,district,-20.30548720,57.48535610
3260,"Port Louis",140,MU,Mauritius,PL,district,-20.16089120,57.50122220
3261,"Rivière du Rempart",140,MU,Mauritius,RR,district,-20.05609830,57.65523890
3249,"Rodrigues Island",140,MU,Mauritius,RO,dependency,-19.72453850,63.42721850
3251,"Saint Brandon Islands",140,MU,Mauritius,CC,dependency,-16.58333300,59.61666700
3257,Savanne,140,MU,Mauritius,SA,district,-20.47395300,57.48535610
3456,Aguascalientes,142,MX,Mexico,AGU,state,21.88525620,-102.29156770
3457,"Baja California",142,MX,Mexico,BCN,state,30.84063380,-115.28375850
3460,"Baja California Sur",142,MX,Mexico,BCS,state,26.04444460,-111.66607250
3475,Campeche,142,MX,Mexico,CAM,state,19.83012510,-90.53490870
3451,Chiapas,142,MX,Mexico,CHP,state,16.75693180,-93.12923530
3447,Chihuahua,142,MX,Mexico,CHH,state,28.63299570,-106.06910040
3473,"Ciudad de México",142,MX,Mexico,CMX,"federal district",19.43260770,-99.13320800
3471,"Coahuila de Zaragoza",142,MX,Mexico,COA,state,27.05867600,-101.70682940
3472,Colima,142,MX,Mexico,COL,state,19.24523420,-103.72408680
3453,Durango,142,MX,Mexico,DUR,state,37.27528000,-107.88006670
3450,"Estado de México",142,MX,Mexico,MEX,state,23.63450100,-102.55278400
3469,Guanajuato,142,MX,Mexico,GUA,state,21.01901450,-101.25735860
3459,Guerrero,142,MX,Mexico,GRO,state,17.43919260,-99.54509740
3470,Hidalgo,142,MX,Mexico,HID,state,26.10035470,-98.26306840
4857,Jalisco,142,MX,Mexico,JAL,state,20.65953820,-103.34943760
3474,"Michoacán de Ocampo",142,MX,Mexico,MIC,state,19.56651920,-101.70682940
3465,Morelos,142,MX,Mexico,MOR,state,18.68130490,-99.10134980
3477,Nayarit,142,MX,Mexico,NAY,state,21.75138440,-104.84546190
3452,"Nuevo León",142,MX,Mexico,NLE,state,25.59217200,-99.99619470
3448,Oaxaca,142,MX,Mexico,OAX,state,17.07318420,-96.72658890
3476,Puebla,142,MX,Mexico,PUE,state,19.04143980,-98.20627270
3455,Querétaro,142,MX,Mexico,QUE,state,20.58879320,-100.38988810
3467,"Quintana Roo",142,MX,Mexico,ROO,state,19.18173930,-88.47913760
3461,"San Luis Potosí",142,MX,Mexico,SLP,state,22.15646990,-100.98554090
3449,Sinaloa,142,MX,Mexico,SIN,state,25.17210910,-107.47951730
3468,Sonora,142,MX,Mexico,SON,state,37.98294960,-120.38217240
3454,Tabasco,142,MX,Mexico,TAB,state,17.84091730,-92.61892730
3463,Tamaulipas,142,MX,Mexico,TAM,state,24.26694000,-98.83627550
3458,Tlaxcala,142,MX,Mexico,TLA,state,19.31815400,-98.23749540
3464,"Veracruz de Ignacio de la Llave",142,MX,Mexico,VER,state,19.17377300,-96.13422410
3466,Yucatán,142,MX,Mexico,YUC,state,20.70987860,-89.09433770
3462,Zacatecas,142,MX,Mexico,ZAC,state,22.77085550,-102.58324260
2580,"Chuuk State",143,FM,Micronesia,TRK,,7.13867590,151.55930650
2583,"Kosrae State",143,FM,Micronesia,KSA,,5.30956180,162.98148770
2581,"Pohnpei State",143,FM,Micronesia,PNI,,6.85412540,158.26238220
2582,"Yap State",143,FM,Micronesia,YAP,,8.67164900,142.84393350
4368,"Anenii Noi District",144,MD,Moldova,AN,,46.87956630,29.23121750
4393,"Bălți Municipality",144,MD,Moldova,BA,,47.75399470,27.91841480
4379,"Basarabeasca District",144,MD,Moldova,BS,,46.42370600,28.89354920
4362,"Bender Municipality",144,MD,Moldova,BD,,46.82275510,29.46201010
4375,"Briceni District",144,MD,Moldova,BR,,48.36320220,27.07503980
4391,"Cahul District",144,MD,Moldova,CA,,45.89394040,28.18902750
4366,"Călărași District",144,MD,Moldova,CL,,47.28694600,28.27453100
4380,"Cantemir District",144,MD,Moldova,CT,,46.27717420,28.20096530
4365,"Căușeni District",144,MD,Moldova,CS,,46.65547150,29.40912220
4373,"Chișinău Municipality",144,MD,Moldova,CU,,47.01045290,28.86381020
4360,"Cimișlia District",144,MD,Moldova,CM,,46.52508510,28.77218350
4390,"Criuleni District",144,MD,Moldova,CR,,47.21361140,29.15575190
4384,"Dondușeni District",144,MD,Moldova,DO,,48.23383050,27.59980870
4392,"Drochia District",144,MD,Moldova,DR,,48.07977880,27.86041140
4383,"Dubăsari District",144,MD,Moldova,DU,,47.26439420,29.15503480
4387,"Edineț District",144,MD,Moldova,ED,,48.16789910,27.29361430
4381,"Fălești District",144,MD,Moldova,FA,,47.56477250,27.72655930
4370,"Florești District",144,MD,Moldova,FL,,47.86678490,28.33918640
4385,Gagauzia,144,MD,Moldova,GA,,46.09794350,28.63846450
4367,"Glodeni District",144,MD,Moldova,GL,,47.77901560,27.51680100
4382,"Hîncești District",144,MD,Moldova,HI,,46.82811470,28.58508890
4369,"Ialoveni District",144,MD,Moldova,IA,,46.86308600,28.82342180
4363,"Nisporeni District",144,MD,Moldova,NI,,47.07513490,28.17681550
4389,"Ocnița District",144,MD,Moldova,OC,,48.41104350,27.47680920
4361,"Orhei District",144,MD,Moldova,OR,,47.38604000,28.83030820
4394,"Rezina District",144,MD,Moldova,RE,,47.71804470,28.88710240
4376,"Rîșcani District",144,MD,Moldova,RI,,47.90701530,27.53749960
4364,"Sîngerei District",144,MD,Moldova,SI,,47.63891340,28.13718160
4388,"Șoldănești District",144,MD,Moldova,SD,,47.81473890,28.78895860
4374,"Soroca District",144,MD,Moldova,SO,,48.15497430,28.28707830
4378,"Ștefan Vodă District",144,MD,Moldova,SV,,46.55404880,29.70224200
4377,"Strășeni District",144,MD,Moldova,ST,,47.14502670,28.61367360
4372,"Taraclia District",144,MD,Moldova,TA,,45.89865100,28.66716440
4371,"Telenești District",144,MD,Moldova,TE,,47.49839620,28.36760190
4395,"Transnistria autonomous territorial unit",144,MD,Moldova,SN,,47.21529720,29.46380540
4386,"Ungheni District",144,MD,Moldova,UN,,47.23057670,27.78926610
4917,"La Colle",145,MC,Monaco,CL,,43.73274650,7.41372760
4918,"La Condamine",145,MC,Monaco,CO,,43.73506650,7.41990600
4919,Moneghetti,145,MC,Monaco,MG,,43.73649270,7.41533830
1973,"Arkhangai Province",146,MN,Mongolia,073,,47.89711010,100.72401650
1969,"Bayan-Ölgii Province",146,MN,Mongolia,071,,48.39832540,89.66259150
1976,"Bayankhongor Province",146,MN,Mongolia,069,,45.15267070,100.10736670
1961,"Bulgan Province",146,MN,Mongolia,067,,48.96909130,102.88317230
1962,"Darkhan-Uul Province",146,MN,Mongolia,037,,49.46484340,105.97459190
1963,"Dornod Province",146,MN,Mongolia,061,,47.46581540,115.39271200
1981,"Dornogovi Province",146,MN,Mongolia,063,,43.96538890,109.17734590
1970,"Dundgovi Province",146,MN,Mongolia,059,,45.58227860,106.76442090
1972,"Govi-Altai Province",146,MN,Mongolia,065,,45.45112270,95.85057660
1978,"Govisümber Province",146,MN,Mongolia,064,,46.47627540,108.55706270
1974,"Khentii Province",146,MN,Mongolia,039,,47.60812090,109.93728560
1964,"Khovd Province",146,MN,Mongolia,043,,47.11296540,92.31107520
1975,"Khövsgöl Province",146,MN,Mongolia,041,,50.22044840,100.32137680
1967,"Ömnögovi Province",146,MN,Mongolia,053,,43.50002400,104.28611160
1966,"Orkhon Province",146,MN,Mongolia,035,,49.00470500,104.30165270
1965,"Övörkhangai Province",146,MN,Mongolia,055,,45.76243920,103.09170320
1980,"Selenge Province",146,MN,Mongolia,049,,50.00592730,106.44341080
1977,"Sükhbaatar Province",146,MN,Mongolia,051,,46.56531630,113.53808360
1968,"Töv Province",146,MN,Mongolia,047,,47.21240560,106.41541000
1971,"Uvs Province",146,MN,Mongolia,046,,49.64497070,93.27365760
1979,"Zavkhan Province",146,MN,Mongolia,057,,48.23881470,96.07030190
23,"Andrijevica Municipality",147,ME,Montenegro,01,,42.73624770,19.78595560
13,"Bar Municipality",147,ME,Montenegro,02,,42.12781190,19.14043800
21,"Berane Municipality",147,ME,Montenegro,03,,42.82572890,19.90205090
25,"Bijelo Polje Municipality",147,ME,Montenegro,04,,43.08465260,19.71154720
30,"Budva Municipality",147,ME,Montenegro,05,,42.31407200,18.83138320
14,"Danilovgrad Municipality",147,ME,Montenegro,07,,42.58357000,19.14043800
24,"Gusinje Municipality",147,ME,Montenegro,22,,42.55634550,19.83060510
31,"Kolašin Municipality",147,ME,Montenegro,09,,42.76019160,19.42591140
26,"Kotor Municipality",147,ME,Montenegro,10,,42.57402610,18.64131450
22,"Mojkovac Municipality",147,ME,Montenegro,11,,42.96880180,19.52110630
17,"Nikšić Municipality",147,ME,Montenegro,12,,42.79971840,18.76009630
28,"Old Royal Capital Cetinje",147,ME,Montenegro,06,,42.39309590,18.91159640
12,"Petnjica Municipality",147,ME,Montenegro,23,,42.93534800,20.02114490
19,"Plav Municipality",147,ME,Montenegro,13,,42.60013370,19.94075410
20,"Pljevlja Municipality",147,ME,Montenegro,14,,43.27233830,19.28315310
16,"Plužine Municipality",147,ME,Montenegro,15,,43.15933840,18.85514840
27,"Podgorica Municipality",147,ME,Montenegro,16,,42.36938340,19.28315310
15,"Rožaje Municipality",147,ME,Montenegro,17,,42.84083890,20.16706280
18,"Šavnik Municipality",147,ME,Montenegro,18,,42.96037560,19.14043800
29,"Tivat Municipality",147,ME,Montenegro,19,,42.42348000,18.71851840
33,"Ulcinj Municipality",147,ME,Montenegro,20,,41.96527950,19.30694320
32,"Žabljak Municipality",147,ME,Montenegro,21,,43.15551520,19.12260180
4928,Agadir-Ida-Ou-Tanane,149,MA,Morocco,AGD,prefecture,30.64620910,-9.83390610
3320,"Al Haouz",149,MA,Morocco,HAO,province,31.29567290,-7.87216000
3267,"Al Hoceïma",149,MA,Morocco,HOC,province,35.24455890,-3.93174680
3266,"Aousserd (EH)",149,MA,Morocco,AOU,province,22.55215380,-14.32973530
3297,"Assa-Zag (EH-partial)",149,MA,Morocco,ASZ,province,28.14023950,-9.72326730
3321,Azilal,149,MA,Morocco,AZI,province,32.00426200,-6.57833870
3272,"Béni Mellal",149,MA,Morocco,BEM,province,32.34244300,-6.37579900
3278,"Béni Mellal-Khénifra",149,MA,Morocco,05,region,32.57191840,-6.06791940
3304,Benslimane,149,MA,Morocco,BES,province,33.61896980,-7.13055360
3285,Berkane,149,MA,Morocco,BER,province,34.88408760,-2.34188700
4929,Berrechid,149,MA,Morocco,BRR,province,33.26025230,-7.59848370
3275,"Boujdour (EH)",149,MA,Morocco,BOD,province,26.12524930,-14.48473470
3270,Boulemane,149,MA,Morocco,BOM,province,33.36251590,-4.73033970
4930,Casablanca,149,MA,Morocco,CAS,prefecture,33.57226780,-7.65703260
3303,Casablanca-Settat,149,MA,Morocco,06,region,33.21608720,-7.43813550
3310,Chefchaouen,149,MA,Morocco,CHE,province,35.01817200,-5.14320680
3274,Chichaoua,149,MA,Morocco,CHI,province,31.53835810,-8.76463880
3302,"Chtouka-Ait Baha",149,MA,Morocco,CHT,province,30.10724220,-9.27855830
3306,"Dakhla-Oued Ed-Dahab (EH)",149,MA,Morocco,12,region,22.73378920,-14.28611160
3290,Drâa-Tafilalet,149,MA,Morocco,08,region,31.14995380,-5.39395510
4931,Driouch,149,MA,Morocco,DRI,province,34.97603200,-3.39644930
3291,"El Hajeb",149,MA,Morocco,HAJ,province,33.68573500,-5.36778440
3280,"El Jadida",149,MA,Morocco,JDI,province,33.23163260,-8.50071160
3309,"El Kelâa des Sraghna",149,MA,Morocco,KES,province,32.05227670,-7.35165580
3299,Errachidia,149,MA,Morocco,ERR,province,31.90512750,-4.72775280
3292,"Es-Semara (EH-partial)",149,MA,Morocco,ESM,province,26.74185600,-11.67836710
3316,Essaouira,149,MA,Morocco,ESI,province,31.50849260,-9.75950410
3300,Fahs-Anjra,149,MA,Morocco,FAH,province,35.76019920,-5.66683060
4932,Fès,149,MA,Morocco,FES,prefecture,34.02395790,-5.03675990
3313,Fès-Meknès,149,MA,Morocco,03,region,34.06252900,-4.72775280
3301,Figuig,149,MA,Morocco,FIG,province,32.10926130,-1.22980600
4933,"Fquih Ben Salah",149,MA,Morocco,FQH,province,32.50016800,-6.71007170
3265,Guelmim,149,MA,Morocco,GUE,province,28.98836590,-10.05274980
3305,"Guelmim-Oued Noun (EH-partial)",149,MA,Morocco,10,region,28.48442810,-10.08072980
4934,Guercif,149,MA,Morocco,GUF,province,34.23450360,-3.38130050
3325,Ifrane,149,MA,Morocco,IFR,province,33.52280620,-5.11095520
3294,"Inezgane-Ait Melloul",149,MA,Morocco,INE,prefecture,30.35090980,-9.38951100
3307,Jerada,149,MA,Morocco,JRA,province,34.30617910,-2.17941360
3308,Kénitra,149,MA,Morocco,KEN,province,34.25405030,-6.58901660
3276,Khémisset,149,MA,Morocco,KHE,province,33.81537040,-6.05733020
3317,Khénifra,149,MA,Morocco,KHN,province,32.93404710,-5.66157100
3326,Khouribga,149,MA,Morocco,KHO,province,32.88602300,-6.92086550
3271,L'Oriental,149,MA,Morocco,02,region,37.06968300,-94.51227700
3293,"Laâyoune (EH)",149,MA,Morocco,LAA,province,27.15003840,-13.19907580
3298,"Laâyoune-Sakia El Hamra (EH-partial)",149,MA,Morocco,11,region,27.86831940,-11.98046130
3268,Larache,149,MA,Morocco,LAR,province,35.17442710,-6.14739640
4936,M’diq-Fnideq,149,MA,Morocco,MDF,prefecture,35.77330190,-5.51433000
4935,Marrakech,149,MA,Morocco,MAR,prefecture,31.63460230,-8.07789320
3288,Marrakesh-Safi,149,MA,Morocco,07,region,31.73308330,-8.13385580
3284,Médiouna,149,MA,Morocco,MED,province,33.45409390,-7.51660200
4937,Meknès,149,MA,Morocco,MEK,prefecture,33.88100000,-5.57303970
4938,Midelt,149,MA,Morocco,MID,province,32.68550790,-4.75017090
4939,Mohammadia,149,MA,Morocco,MOH,prefecture,33.68737490,-7.42391420
3315,"Moulay Yacoub",149,MA,Morocco,MOU,province,34.08744790,-5.17840190
3281,Nador,149,MA,Morocco,NAD,province,34.91719260,-2.85771050
3287,Nouaceur,149,MA,Morocco,NOU,province,33.36703930,-7.57325370
3269,Ouarzazate,149,MA,Morocco,OUA,province,30.93354360,-6.93701600
3319,"Oued Ed-Dahab (EH)",149,MA,Morocco,OUD,province,22.73378920,-14.28611160
4941,Ouezzane,149,MA,Morocco,OUZ,province,34.80634500,-5.59145050
4940,Oujda-Angad,149,MA,Morocco,OUJ,prefecture,34.68375040,-2.29932390
4942,Rabat,149,MA,Morocco,RAB,prefecture,33.96919900,-6.92730290
4927,Rabat-Salé-Kénitra,149,MA,Morocco,04,region,34.07686400,-7.34544760
4943,Rehamna,149,MA,Morocco,REH,province,32.20329050,-8.56896710
3311,Safi,149,MA,Morocco,SAF,province,32.29898720,-9.10134980
4944,Salé,149,MA,Morocco,SAL,prefecture,34.03775700,-6.84270730
3289,Sefrou,149,MA,Morocco,SEF,province,33.83052440,-4.83531540
3282,Settat,149,MA,Morocco,SET,province,32.99242420,-7.62226650
4945,"Sidi Bennour",149,MA,Morocco,SIB,province,32.64926020,-8.44714530
4946,"Sidi Ifni",149,MA,Morocco,SIF,province,29.36657970,-10.21084850
3279,"Sidi Kacem",149,MA,Morocco,SIK,province,34.22601720,-5.71291640
4952,"Sidi Slimane",149,MA,Morocco,SIL,province,34.27378280,-5.98059720
4947,Skhirate-Témara,149,MA,Morocco,SKH,prefecture,33.76224250,-7.04190520
3295,Souss-Massa,149,MA,Morocco,09,region,30.27506110,-8.13385580
3286,"Tan-Tan (EH-partial)",149,MA,Morocco,TNT,province,28.03012000,-11.16173560
4950,Tanger-Assilah,149,MA,Morocco,TNG,prefecture,35.76325390,-5.90450980
3324,"Tanger-Tétouan-Al Hoceïma",149,MA,Morocco,01,region,35.26295580,-5.56172790
3323,Taounate,149,MA,Morocco,TAO,province,34.53691700,-4.63986930
3322,Taourirt,149,MA,Morocco,TAI,province,34.21259800,-2.69838680
4948,"Tarfaya (EH-partial)",149,MA,Morocco,TAF,province,27.93777010,-12.92940630
3314,Taroudannt,149,MA,Morocco,TAR,province,30.47271260,-8.87487650
3312,Tata,149,MA,Morocco,TAT,province,29.75087700,-7.97563430
3296,Taza,149,MA,Morocco,TAZ,province,34.27889530,-3.58126920
3318,Tétouan,149,MA,Morocco,TET,province,35.58889950,-5.36255160
4949,Tinghir,149,MA,Morocco,TIN,province,31.48507940,-6.20192980
3277,Tiznit,149,MA,Morocco,TIZ,province,29.69339200,-9.73215700
4951,Youssoufia,149,MA,Morocco,YUS,province,32.02006790,-8.86926480
3283,Zagora,149,MA,Morocco,ZAG,province,30.57860930,-5.89871390
3327,"Cabo Delgado Province",150,MZ,Mozambique,P,,-12.33354740,39.32062410
3329,"Gaza Province",150,MZ,Mozambique,G,,-23.02219280,32.71813750
3330,"Inhambane Province",150,MZ,Mozambique,I,,-22.85279970,34.55087580
3337,"Manica Province",150,MZ,Mozambique,B,,-19.50597870,33.43835300
3335,Maputo,150,MZ,Mozambique,MPM,,-25.96924800,32.57317460
3332,"Maputo Province",150,MZ,Mozambique,L,,-25.25698760,32.53727410
3336,"Nampula Province",150,MZ,Mozambique,N,,-14.76049310,39.32062410
3333,"Niassa Province",150,MZ,Mozambique,A,,-12.78262020,36.60939260
3331,"Sofala Province",150,MZ,Mozambique,S,,-19.20390730,34.86241660
3334,"Tete Province",150,MZ,Mozambique,T,,-15.65960560,32.71813750
3328,"Zambezia Province",150,MZ,Mozambique,Q,,-16.56389870,36.60939260
2142,"Ayeyarwady Region",151,MM,Myanmar,07,,17.03421250,95.22666750
2141,Bago,151,MM,Myanmar,02,,17.32207110,96.46632860
2137,"Chin State",151,MM,Myanmar,14,,22.00869780,93.58126920
2143,"Kachin State",151,MM,Myanmar,11,,25.85090400,97.43813550
2144,"Kayah State",151,MM,Myanmar,12,,19.23420610,97.26528580
2133,"Kayin State",151,MM,Myanmar,13,,16.94593460,97.95928630
2136,"Magway Region",151,MM,Myanmar,03,,19.88713860,94.72775280
2134,"Mandalay Region",151,MM,Myanmar,04,,21.56190580,95.89871390
2147,"Mon State",151,MM,Myanmar,15,,16.30031330,97.69822720
2146,"Naypyidaw Union Territory",151,MM,Myanmar,18,,19.93862450,96.15269850
2138,"Rakhine State",151,MM,Myanmar,16,,20.10408180,93.58126920
2145,"Sagaing Region",151,MM,Myanmar,01,,24.42838100,95.39395510
2139,"Shan State",151,MM,Myanmar,17,,22.03619850,98.13385580
2140,"Tanintharyi Region",151,MM,Myanmar,05,,12.47068760,99.01289260
2135,"Yangon Region",151,MM,Myanmar,06,,16.91434880,96.15269850
43,"Erongo Region",152,NA,Namibia,ER,,-22.25656820,15.40680790
38,"Hardap Region",152,NA,Namibia,HA,,-24.23101340,17.66888700
45,"Karas Region",152,NA,Namibia,KA,,-26.84296450,17.29028390
36,"Kavango East Region",152,NA,Namibia,KE,,-18.27104800,18.42760470
35,"Kavango West Region",152,NA,Namibia,KW,,-18.27104800,18.42760470
44,"Khomas Region",152,NA,Namibia,KH,,-22.63778540,17.10119310
34,"Kunene Region",152,NA,Namibia,KU,,-19.40863170,13.91439900
40,"Ohangwena Region",152,NA,Namibia,OW,,-17.59792910,16.81783770
41,"Omaheke Region",152,NA,Namibia,OH,,-21.84666510,19.18800470
39,"Omusati Region",152,NA,Namibia,OS,,-18.40702940,14.84546190
37,"Oshana Region",152,NA,Namibia,ON,,-18.43050640,15.68817880
42,"Oshikoto Region",152,NA,Namibia,OT,,-18.41525750,16.91225100
46,"Otjozondjupa Region",152,NA,Namibia,OD,,-20.54869160,17.66888700
47,"Zambezi Region",152,NA,Namibia,CA,,-17.81934190,23.95364660
4656,"Aiwo District",153,NR,Nauru,01,,-0.53400120,166.91388730
4658,"Anabar District",153,NR,Nauru,02,,-0.51335170,166.94846240
4667,"Anetan District",153,NR,Nauru,03,,-0.50643430,166.94270060
4663,"Anibare District",153,NR,Nauru,04,,-0.52947580,166.95134320
4660,"Baiti District",153,NR,Nauru,05,,-0.51043100,166.92757440
4665,"Boe District",153,NR,Nauru,06,,39.07327760,-94.57104980
4662,"Buada District",153,NR,Nauru,07,,-0.53287770,166.92685410
4666,"Denigomodu District",153,NR,Nauru,08,,-0.52479640,166.91676890
4654,"Ewa District",153,NR,Nauru,09,,-0.50872410,166.93693840
4661,"Ijuw District",153,NR,Nauru,10,,-0.52027670,166.95710460
4657,"Meneng District",153,NR,Nauru,11,,-0.54672400,166.93837900
4659,"Nibok District",153,NR,Nauru,12,,-0.51962080,166.91893010
4655,"Uaboe District",153,NR,Nauru,13,,-0.52022220,166.93117610
4664,"Yaren District",153,NR,Nauru,14,,-0.54668570,166.92109130
2082,"Bagmati Zone",154,NP,Nepal,BA,,28.03675770,85.43755740
2071,"Bheri Zone",154,NP,Nepal,BH,,28.51745600,81.77870210
2073,"Central Region",154,NP,Nepal,1,,,
2080,"Dhaulagiri Zone",154,NP,Nepal,DH,,28.61117600,83.50702030
2069,"Eastern Development Region",154,NP,Nepal,4,,27.33090720,87.06242610
2068,"Far-Western Development Region",154,NP,Nepal,5,,29.29878710,80.98710740
2081,"Gandaki Zone",154,NP,Nepal,GA,,28.37320370,84.43827210
2076,"Janakpur Zone",154,NP,Nepal,JA,,27.21108990,86.01215730
2079,"Karnali Zone",154,NP,Nepal,KA,,29.38625550,82.38857830
2072,"Kosi Zone",154,NP,Nepal,KO,,27.05365240,87.30161320
2074,"Lumbini Zone",154,NP,Nepal,LU,,27.45000000,83.25000000
2083,"Mahakali Zone",154,NP,Nepal,MA,,29.36010790,80.54384500
2070,"Mechi Zone",154,NP,Nepal,ME,,26.87600070,87.93348030
2066,"Mid-Western Region",154,NP,Nepal,2,,38.41118410,-90.38320980
2075,"Narayani Zone",154,NP,Nepal,NA,,27.36117660,84.85679320
2077,"Rapti Zone",154,NP,Nepal,RA,,28.27434700,82.38857830
2084,"Sagarmatha Zone",154,NP,Nepal,SA,,27.32382630,86.74163740
2078,"Seti Zone",154,NP,Nepal,SE,,29.69054270,81.33994140
2067,"Western Region",154,NP,Nepal,3,,,
2624,Bonaire,156,NL,Netherlands,BQ1,"special municipality",12.20189020,-68.26238220
2613,Drenthe,156,NL,Netherlands,DR,province,52.94760120,6.62305860
2619,Flevoland,156,NL,Netherlands,FL,province,52.52797810,5.59535080
2622,Friesland,156,NL,Netherlands,FR,province,53.16416420,5.78175420
2611,Gelderland,156,NL,Netherlands,GE,province,52.04515500,5.87182350
2617,Groningen,156,NL,Netherlands,GR,province,53.21938350,6.56650170
2615,Limburg,156,NL,Netherlands,LI,province,51.44272380,6.06087260
2623,"North Brabant",156,NL,Netherlands,NB,province,51.48265370,5.23216870
2612,"North Holland",156,NL,Netherlands,NH,province,52.52058690,4.78847400
2618,Overijssel,156,NL,Netherlands,OV,province,52.43878140,6.50164110
2621,Saba,156,NL,Netherlands,BQ2,"special municipality",17.63546420,-63.23267630
2616,"Sint Eustatius",156,NL,Netherlands,BQ3,"special municipality",17.48903060,-62.97355500
2614,"South Holland",156,NL,Netherlands,ZH,province,51.99667920,4.55973970
2610,Utrecht,156,NL,Netherlands,UT,province,52.09073740,5.12142010
2620,Zeeland,156,NL,Netherlands,ZE,province,51.49403090,3.84968150
5227,"Loyalty Islands Province",157,NC,"New Caledonia",03,province,-20.96670000,167.23330000
5226,"North Province",157,NC,"New Caledonia",02,province,-22.27580000,166.45800000
5225,"South Province",157,NC,"New Caledonia",01,province,-22.27580000,166.45800000
4072,"Auckland Region",158,NZ,"New Zealand",AUK,,-36.66753280,174.77333250
4074,"Bay of Plenty Region",158,NZ,"New Zealand",BOP,,-37.42339170,176.74163740
4066,"Canterbury Region",158,NZ,"New Zealand",CAN,,-43.75422750,171.16372450
4067,"Chatham Islands",158,NZ,"New Zealand",CIT,,-44.00575230,-176.54006740
4068,"Gisborne District",158,NZ,"New Zealand",GIS,,-38.13581740,178.32393090
4075,"Hawke's Bay Region",158,NZ,"New Zealand",HKB,,-39.60165970,176.58044730
4060,"Manawatu-Wanganui Region",158,NZ,"New Zealand",MWT,,-39.72733560,175.43755740
4063,"Marlborough Region",158,NZ,"New Zealand",MBH,,-41.59168830,173.76240530
4070,"Nelson Region",158,NZ,"New Zealand",NSN,,-41.29853970,173.24414910
4059,"Northland Region",158,NZ,"New Zealand",NTL,,-35.41361720,173.93208060
4062,"Otago Region",158,NZ,"New Zealand",OTA,,-45.47906710,170.15475670
4071,"Southland Region",158,NZ,"New Zealand",STL,,-45.84891590,167.67553870
4069,"Taranaki Region",158,NZ,"New Zealand",TKI,,-39.35381490,174.43827210
4073,"Tasman District",158,NZ,"New Zealand",TAS,,-41.45711840,172.82097400
4061,"Waikato Region",158,NZ,"New Zealand",WKO,,-37.61908620,175.02334600
4065,"Wellington Region",158,NZ,"New Zealand",WGN,,-41.02993230,175.43755740
4064,"West Coast Region",158,NZ,"New Zealand",WTC,,62.41136340,-149.07297140
946,Boaco,159,NI,Nicaragua,BO,department,12.46928400,-85.66146820
950,Carazo,159,NI,Nicaragua,CA,department,11.72747290,-86.21584970
954,Chinandega,159,NI,Nicaragua,CI,department,12.88200620,-87.14228950
940,Chontales,159,NI,Nicaragua,CO,department,11.93947170,-85.18940450
945,Estelí,159,NI,Nicaragua,ES,department,13.08511390,-86.36301970
943,Granada,159,NI,Nicaragua,GR,department,11.93440730,-85.95600050
955,Jinotega,159,NI,Nicaragua,JI,department,13.08839070,-85.99939970
944,León,159,NI,Nicaragua,LE,department,12.50920370,-86.66110830
948,Madriz,159,NI,Nicaragua,MD,department,13.47260050,-86.45920910
941,Managua,159,NI,Nicaragua,MN,department,12.13916990,-86.33767610
953,Masaya,159,NI,Nicaragua,MS,department,11.97593280,-86.07334980
947,Matagalpa,159,NI,Nicaragua,MT,department,12.94984360,-85.43755740
951,"North Caribbean Coast",159,NI,Nicaragua,AN,"autonomous region",13.83944560,-83.93208060
4964,"Nueva Segovia",159,NI,Nicaragua,NS,department,13.76570610,-86.53700390
949,"Río San Juan",159,NI,Nicaragua,SJ,department,11.47816100,-84.77333250
942,Rivas,159,NI,Nicaragua,RI,department,11.40234900,-85.68457800
952,"South Caribbean Coast",159,NI,Nicaragua,AS,"autonomous region",12.19185020,-84.10128610
71,"Agadez Region",160,NE,Niger,1,,20.66707520,12.07182810
72,"Diffa Region",160,NE,Niger,2,,13.67686470,12.71351210
68,"Dosso Region",160,NE,Niger,3,,13.15139470,3.41955270
70,"Maradi Region",160,NE,Niger,4,,13.80180740,7.43813550
73,"Tahoua Region",160,NE,Niger,5,,16.09025430,5.39395510
67,"Tillabéri Region",160,NE,Niger,6,,14.64895250,2.14502450
69,"Zinder Region",160,NE,Niger,7,,15.17188810,10.26001250
303,Abia,161,NG,Nigeria,AB,state,5.45273540,7.52484140
293,"Abuja Federal Capital Territory",161,NG,Nigeria,FC,"capital territory",8.89406910,7.18604020
320,Adamawa,161,NG,Nigeria,AD,state,9.32647510,12.39838530
304,"Akwa Ibom",161,NG,Nigeria,AK,state,4.90573710,7.85366750
315,Anambra,161,NG,Nigeria,AN,state,6.22089970,6.93695590
312,Bauchi,161,NG,Nigeria,BA,state,10.77606240,9.99919430
305,Bayelsa,161,NG,Nigeria,BY,state,4.77190710,6.06985260
291,Benue,161,NG,Nigeria,BE,state,7.33690240,8.74036870
307,Borno,161,NG,Nigeria,BO,state,11.88463560,13.15196650
314,"Cross River",161,NG,Nigeria,CR,state,5.87017240,8.59880140
316,Delta,161,NG,Nigeria,DE,state,33.74537840,-90.73545080
311,Ebonyi,161,NG,Nigeria,EB,state,6.26492320,8.01373020
318,Edo,161,NG,Nigeria,ED,state,6.63418310,5.93040560
309,Ekiti,161,NG,Nigeria,EK,state,7.71898620,5.31095050
289,Enugu,161,NG,Nigeria,EN,state,6.53635300,7.43561940
310,Gombe,161,NG,Nigeria,GO,state,10.36377950,11.19275870
308,Imo,161,NG,Nigeria,IM,state,5.57201220,7.05882190
288,Jigawa,161,NG,Nigeria,JI,state,12.22801200,9.56158670
294,Kaduna,161,NG,Nigeria,KD,state,10.37640060,7.70945370
300,Kano,161,NG,Nigeria,KN,state,11.74706980,8.52471070
313,Katsina,161,NG,Nigeria,KT,state,12.37967070,7.63057480
290,Kebbi,161,NG,Nigeria,KE,state,11.49420030,4.23333550
298,Kogi,161,NG,Nigeria,KO,state,7.73373250,6.69058360
295,Kwara,161,NG,Nigeria,KW,state,8.96689610,4.38740510
306,Lagos,161,NG,Nigeria,LA,state,6.52437930,3.37920570
301,Nasarawa,161,NG,Nigeria,NA,state,8.49979080,8.19969370
317,Niger,161,NG,Nigeria,NI,state,9.93092240,5.59832100
323,Ogun,161,NG,Nigeria,OG,state,6.99797470,3.47373780
321,Ondo,161,NG,Nigeria,ON,state,6.91486820,5.14781440
322,Osun,161,NG,Nigeria,OS,state,7.56289640,4.51995930
296,Oyo,161,NG,Nigeria,OY,state,8.15738090,3.61465340
302,Plateau,161,NG,Nigeria,PL,state,9.21820930,9.51794880
4926,Rivers,161,NG,Nigeria,RI,state,5.02134200,6.43760220
292,Sokoto,161,NG,Nigeria,SO,state,13.05331430,5.32227220
319,Taraba,161,NG,Nigeria,TA,state,7.99936160,10.77398630
297,Yobe,161,NG,Nigeria,YO,state,12.29387600,11.43904110
299,Zamfara,161,NG,Nigeria,ZA,state,12.12218050,6.22358190
3998,"Chagang Province",115,KP,"North Korea",04,,40.72028090,126.56211370
3999,"Kangwon Province",115,KP,"North Korea",07,,38.84323930,127.55970670
3995,"North Hamgyong Province",115,KP,"North Korea",09,,41.81487580,129.45819550
4004,"North Hwanghae Province",115,KP,"North Korea",06,,38.37860850,126.43643630
4002,"North Pyongan Province",115,KP,"North Korea",03,,39.92556180,125.39280250
4005,Pyongyang,115,KP,"North Korea",01,,39.03921930,125.76252410
4001,Rason,115,KP,"North Korea",13,,42.25690630,130.29771860
3996,"Ryanggang Province",115,KP,"North Korea",10,,41.23189210,128.50763590
4000,"South Hamgyong Province",115,KP,"North Korea",08,,40.37253390,128.29888400
4003,"South Hwanghae Province",115,KP,"North Korea",05,,38.20072150,125.47819260
3997,"South Pyongan Province",115,KP,"North Korea",02,,39.35391780,126.16827100
703,"Aerodrom Municipality",129,MK,"North Macedonia",01,,41.94643630,21.49317130
656,"Aračinovo Municipality",129,MK,"North Macedonia",02,,42.02473810,21.57664070
716,"Berovo Municipality",129,MK,"North Macedonia",03,,41.66619290,22.76288300
679,"Bitola Municipality",129,MK,"North Macedonia",04,,41.03633020,21.33219740
649,"Bogdanci Municipality",129,MK,"North Macedonia",05,,41.18696160,22.59602680
721,"Bogovinje Municipality",129,MK,"North Macedonia",06,,41.92363710,20.91638870
652,"Bosilovo Municipality",129,MK,"North Macedonia",07,,41.49048640,22.78671740
660,"Brvenica Municipality",129,MK,"North Macedonia",08,,41.96814820,20.98195860
694,"Butel Municipality",129,MK,"North Macedonia",09,,42.08950680,21.46336100
704,"Čair Municipality",129,MK,"North Macedonia",79,,41.99303550,21.43653180
676,"Čaška Municipality",129,MK,"North Macedonia",80,,41.64743800,21.69141150
702,"Centar Municipality",129,MK,"North Macedonia",77,,41.96989340,21.42162670
720,"Centar Župa Municipality",129,MK,"North Macedonia",78,,41.46522590,20.59305480
644,"Češinovo-Obleševo Municipality",129,MK,"North Macedonia",81,,41.86393160,22.26224600
715,"Čučer-Sandevo Municipality",129,MK,"North Macedonia",82,,42.14839460,21.40374070
645,"Debarca Municipality",129,MK,"North Macedonia",22,,41.35840770,20.85529190
695,"Delčevo Municipality",129,MK,"North Macedonia",23,,41.96843870,22.76288300
687,"Demir Hisar Municipality",129,MK,"North Macedonia",25,,41.22708300,21.14142260
655,"Demir Kapija Municipality",129,MK,"North Macedonia",24,,41.37955380,22.21455710
697,"Dojran Municipality",129,MK,"North Macedonia",26,,41.24366720,22.69137640
675,"Dolneni Municipality",129,MK,"North Macedonia",27,,41.46409350,21.40374070
657,"Drugovo Municipality",129,MK,"North Macedonia",28,,41.44081530,20.92682010
707,"Gazi Baba Municipality",129,MK,"North Macedonia",17,,42.01629610,21.49913340
648,"Gevgelija Municipality",129,MK,"North Macedonia",18,,41.21186060,22.38146240
722,"Gjorče Petrov Municipality",129,MK,"North Macedonia",29,,42.06063740,21.32027360
693,"Gostivar Municipality",129,MK,"North Macedonia",19,,41.80255410,20.90893780
708,"Gradsko Municipality",129,MK,"North Macedonia",20,,41.59916080,21.88070640
684,"Greater Skopje",129,MK,"North Macedonia",85,,41.99812940,21.42543550
690,"Ilinden Municipality",129,MK,"North Macedonia",34,,41.99574430,21.56769750
678,"Jegunovce Municipality",129,MK,"North Macedonia",35,,42.07407200,21.12204780
674,Karbinci,129,MK,"North Macedonia",37,,41.81801590,22.23247580
681,"Karpoš Municipality",129,MK,"North Macedonia",38,,41.97096610,21.39181680
713,"Kavadarci Municipality",129,MK,"North Macedonia",36,,41.28900680,21.99994350
688,"Kičevo Municipality",129,MK,"North Macedonia",40,,41.51291120,20.95250650
686,"Kisela Voda Municipality",129,MK,"North Macedonia",39,,41.92748000,21.49317130
723,"Kočani Municipality",129,MK,"North Macedonia",42,,41.98583740,22.40530460
665,"Konče Municipality",129,MK,"North Macedonia",41,,41.51710110,22.38146240
641,"Kratovo Municipality",129,MK,"North Macedonia",43,,42.05371410,22.07148350
677,"Kriva Palanka Municipality",129,MK,"North Macedonia",44,,42.20584540,22.33079650
647,"Krivogaštani Municipality",129,MK,"North Macedonia",45,,41.30823060,21.36796890
714,"Kruševo Municipality",129,MK,"North Macedonia",46,,41.37693310,21.26065540
683,"Kumanovo Municipality",129,MK,"North Macedonia",47,,42.07326130,21.78531430
659,"Lipkovo Municipality",129,MK,"North Macedonia",48,,42.20066260,21.61837550
705,"Lozovo Municipality",129,MK,"North Macedonia",49,,41.78181390,21.90008270
701,"Makedonska Kamenica Municipality",129,MK,"North Macedonia",51,,42.06946040,22.54834900
692,"Makedonski Brod Municipality",129,MK,"North Macedonia",52,,41.51330880,21.21743290
669,"Mavrovo and Rostuša Municipality",129,MK,"North Macedonia",50,,41.60924270,20.60124880
653,"Mogila Municipality",129,MK,"North Macedonia",53,,41.14796450,21.45143690
664,"Negotino Municipality",129,MK,"North Macedonia",54,,41.49899850,22.09532970
696,"Novaci Municipality",129,MK,"North Macedonia",55,,41.04426610,21.45888940
718,"Novo Selo Municipality",129,MK,"North Macedonia",56,,41.43255800,22.88204890
699,"Ohrid Municipality",129,MK,"North Macedonia",58,,41.06820880,20.75992660
682,"Oslomej Municipality",129,MK,"North Macedonia",57,,41.57583910,21.02219600
685,"Pehčevo Municipality",129,MK,"North Macedonia",60,,41.77371320,22.88204890
698,"Petrovec Municipality",129,MK,"North Macedonia",59,,41.90298970,21.68992100
670,"Plasnica Municipality",129,MK,"North Macedonia",61,,41.45463490,21.10565390
666,"Prilep Municipality",129,MK,"North Macedonia",62,,41.26931420,21.71376940
646,"Probištip Municipality",129,MK,"North Macedonia",63,,41.95891460,22.16686700
709,"Radoviš Municipality",129,MK,"North Macedonia",64,,41.64955310,22.47682870
717,"Rankovce Municipality",129,MK,"North Macedonia",65,,42.18081410,22.09532970
712,"Resen Municipality",129,MK,"North Macedonia",66,,40.93680930,21.04604070
691,"Rosoman Municipality",129,MK,"North Macedonia",67,,41.48480060,21.88070640
667,"Saraj Municipality",129,MK,"North Macedonia",68,,41.98694960,21.26065540
719,"Sopište Municipality",129,MK,"North Macedonia",70,,41.86384920,21.30834990
643,"Staro Nagoričane Municipality",129,MK,"North Macedonia",71,,42.21916920,21.90455410
661,"Štip Municipality",129,MK,"North Macedonia",83,,41.70792970,22.19071220
700,"Struga Municipality",129,MK,"North Macedonia",72,,41.31737440,20.66456830
710,"Strumica Municipality",129,MK,"North Macedonia",73,,41.43780040,22.64274280
711,"Studeničani Municipality",129,MK,"North Macedonia",74,,41.92256390,21.53639650
680,"Šuto Orizari Municipality",129,MK,"North Macedonia",84,,42.02904160,21.40970270
640,"Sveti Nikole Municipality",129,MK,"North Macedonia",69,,41.89803120,21.99994350
654,"Tearce Municipality",129,MK,"North Macedonia",75,,42.07775110,21.05349230
663,"Tetovo Municipality",129,MK,"North Macedonia",76,,42.02748600,20.95066360
671,"Valandovo Municipality",129,MK,"North Macedonia",10,,41.32119090,22.50066930
658,"Vasilevo Municipality",129,MK,"North Macedonia",11,,41.47416990,22.64221280
651,"Veles Municipality",129,MK,"North Macedonia",13,,41.72744260,21.71376940
662,"Vevčani Municipality",129,MK,"North Macedonia",12,,41.24075430,20.59156490
672,"Vinica Municipality",129,MK,"North Macedonia",14,,41.85710200,22.57218810
650,"Vraneštica Municipality",129,MK,"North Macedonia",15,,41.48290870,21.05796320
689,"Vrapčište Municipality",129,MK,"North Macedonia",16,,41.87911600,20.83145000
642,"Zajas Municipality",129,MK,"North Macedonia",31,,41.60303280,20.87913430
706,"Zelenikovo Municipality",129,MK,"North Macedonia",32,,41.87338120,21.60272500
668,"Želino Municipality",129,MK,"North Macedonia",30,,41.90065310,21.11757670
673,"Zrnovci Municipality",129,MK,"North Macedonia",33,,41.82282210,22.41722560
1014,Agder,165,NO,Norway,42,county,58.74069340,6.75315210
1009,Innlandet,165,NO,Norway,34,county,61.19357870,5.50832660
1026,"Jan Mayen",165,NO,Norway,22,"arctic region",71.03181800,-8.29203460
1020,"Møre og Romsdal",165,NO,Norway,15,county,62.84068330,7.00714300
1025,Nordland,165,NO,Norway,18,county,67.69305800,12.70739360
1007,Oslo,165,NO,Norway,03,county,59.91386880,10.75224540
1021,Rogaland,165,NO,Norway,11,county,59.14895440,6.01434320
1013,Svalbard,165,NO,Norway,21,"arctic region",77.87497250,20.97518210
1015,"Troms og Finnmark",165,NO,Norway,54,county,69.77890670,18.99401840
1006,Trøndelag,165,NO,Norway,50,county,63.54201250,10.93692670
1024,"Vestfold og Telemark",165,NO,Norway,38,county,59.41174820,7.76471750
1018,Vestland,165,NO,Norway,46,county,60.90694420,3.96270810
1011,Viken,165,NO,Norway,30,county,59.96530050,7.45051440
3058,"Ad Dakhiliyah",166,OM,Oman,DA,,22.85887580,57.53943560
3047,"Ad Dhahirah",166,OM,Oman,ZA,,23.21616740,56.49074440
3048,"Al Batinah North",166,OM,Oman,BS,,24.34198460,56.72989040
3050,"Al Batinah Region",166,OM,Oman,BA,,24.34198460,56.72989040
3049,"Al Batinah South",166,OM,Oman,BJ,,23.43149030,57.42397960
3059,"Al Buraimi",166,OM,Oman,BU,,24.16714130,56.11422530
3056,"Al Wusta",166,OM,Oman,WU,,19.95710780,56.27568460
3053,"Ash Sharqiyah North",166,OM,Oman,SS,,22.71411960,58.53080640
3051,"Ash Sharqiyah Region",166,OM,Oman,SH,,22.71411960,58.53080640
3054,"Ash Sharqiyah South",166,OM,Oman,SJ,,22.01582490,59.32519220
3057,Dhofar,166,OM,Oman,ZU,,17.03221210,54.14252140
3052,Musandam,166,OM,Oman,MU,,26.19861440,56.24609490
3055,Muscat,166,OM,Oman,MA,,23.58803070,58.38287170
3172,"Azad Kashmir",167,PK,Pakistan,JK,,33.92590550,73.78103340
3174,Balochistan,167,PK,Pakistan,BA,,28.49073320,65.09577920
3173,"Federally Administered Tribal Areas",167,PK,Pakistan,TA,,32.66747600,69.85974060
3170,Gilgit-Baltistan,167,PK,Pakistan,GB,,35.80256670,74.98318080
3169,"Islamabad Capital Territory",167,PK,Pakistan,IS,,33.72049970,73.04052770
3171,"Khyber Pakhtunkhwa",167,PK,Pakistan,KP,,34.95262050,72.33111300
3176,Punjab,167,PK,Pakistan,PB,,31.14713050,75.34121790
3175,Sindh,167,PK,Pakistan,SD,,25.89430180,68.52471490
4540,Aimeliik,168,PW,Palau,002,,7.44558590,134.50308780
4528,Airai,168,PW,Palau,004,,7.39661180,134.56902250
4538,Angaur,168,PW,Palau,010,,6.90922300,134.13879340
4529,Hatohobei,168,PW,Palau,050,,3.00706580,131.12377810
4539,Kayangel,168,PW,Palau,100,,8.07000000,134.70277800
4532,Koror,168,PW,Palau,150,,7.33756460,134.48894690
4530,Melekeok,168,PW,Palau,212,,7.51502860,134.59725180
4537,Ngaraard,168,PW,Palau,214,,7.60794000,134.63486450
4533,Ngarchelong,168,PW,Palau,218,,7.71054690,134.63016460
4527,Ngardmau,168,PW,Palau,222,,7.58504860,134.55960890
4531,Ngatpang,168,PW,Palau,224,,7.47109940,134.52664660
4536,Ngchesar,168,PW,Palau,226,,7.45232800,134.57843420
4541,Ngeremlengui,168,PW,Palau,227,,7.51983970,134.55960890
4534,Ngiwal,168,PW,Palau,228,,7.56147640,134.61606190
4526,Peleliu,168,PW,Palau,350,,7.00229060,134.24316280
4535,Sonsorol,168,PW,Palau,370,,5.32681190,132.22391170
5118,Bethlehem,169,PS,"Palestinian Territory Occupied",BTH,governorate,31.70539960,35.19368770
5119,"Deir El Balah",169,PS,"Palestinian Territory Occupied",DEB,governorate,,
5120,Gaza,169,PS,"Palestinian Territory Occupied",GZA,governorate,,
5121,Hebron,169,PS,"Palestinian Territory Occupied",HBN,governorate,31.53260010,35.06394750
5122,Jenin,169,PS,"Palestinian Territory Occupied",JEN,governorate,,
5123,"Jericho and Al Aghwar",169,PS,"Palestinian Territory Occupied",JRH,governorate,,
5124,Jerusalem,169,PS,"Palestinian Territory Occupied",JEM,governorate,,
5125,"Khan Yunis",169,PS,"Palestinian Territory Occupied",KYS,governorate,,
5126,Nablus,169,PS,"Palestinian Territory Occupied",NBS,governorate,,
5127,"North Gaza",169,PS,"Palestinian Territory Occupied",NGZ,governorate,,
5128,Qalqilya,169,PS,"Palestinian Territory Occupied",QQA,governorate,,
5129,Rafah,169,PS,"Palestinian Territory Occupied",RFH,governorate,,
5130,Ramallah,169,PS,"Palestinian Territory Occupied",RBH,governorate,,
5131,Salfit,169,PS,"Palestinian Territory Occupied",SLT,governorate,,
5132,Tubas,169,PS,"Palestinian Territory Occupied",TBS,governorate,,
5133,Tulkarm,169,PS,"Palestinian Territory Occupied",TKM,governorate,,
1393,"Bocas del Toro Province",170,PA,Panama,1,,9.41655210,-82.52077870
1397,"Chiriquí Province",170,PA,Panama,4,,8.58489800,-82.38857830
1387,"Coclé Province",170,PA,Panama,2,,8.62660680,-80.36586500
1386,"Colón Province",170,PA,Panama,3,,9.18519890,-80.05349230
1385,"Darién Province",170,PA,Panama,5,,7.86817130,-77.83672820
1396,"Emberá-Wounaan Comarca",170,PA,Panama,EM,,8.37669830,-77.65361250
1388,"Guna Yala",170,PA,Panama,KY,,9.23443950,-78.19262500
1389,"Herrera Province",170,PA,Panama,6,,7.77042820,-80.72144170
1390,"Los Santos Province",170,PA,Panama,7,,7.59093020,-80.36586500
1391,"Ngöbe-Buglé Comarca",170,PA,Panama,NB,,8.65958330,-81.77870210
1394,"Panamá Oeste Province",170,PA,Panama,10,,9.11967510,-79.29021330
1395,"Panamá Province",170,PA,Panama,8,,9.11967510,-79.29021330
1392,"Veraguas Province",170,PA,Panama,9,,8.12310330,-81.07546570
4831,Bougainville,171,PG,"Papua New Guinea",NSB,,-6.37539190,155.38071010
4847,"Central Province",171,PG,"Papua New Guinea",CPM,,,
4846,"Chimbu Province",171,PG,"Papua New Guinea",CPK,,-6.30876820,144.87312190
4834,"East New Britain",171,PG,"Papua New Guinea",EBR,,-4.61289430,151.88773210
4845,"Eastern Highlands Province",171,PG,"Papua New Guinea",EHG,,-6.58616740,145.66896360
4848,"Enga Province",171,PG,"Papua New Guinea",EPW,,-5.30058490,143.56356370
4839,Gulf,171,PG,"Papua New Guinea",GPK,,37.05483150,-94.43704190
4833,Hela,171,PG,"Papua New Guinea",HLA,,42.33295160,-83.04826180
4832,"Jiwaka Province",171,PG,"Papua New Guinea",JWK,,-5.86911540,144.69727740
4843,"Madang Province",171,PG,"Papua New Guinea",MPM,,-4.98497330,145.13758340
4842,"Manus Province",171,PG,"Papua New Guinea",MRL,,-2.09411690,146.87609510
4849,"Milne Bay Province",171,PG,"Papua New Guinea",MBA,,-9.52214510,150.67496530
4835,"Morobe Province",171,PG,"Papua New Guinea",MPL,,-6.80137370,146.56164700
4841,"New Ireland Province",171,PG,"Papua New Guinea",NIK,,-4.28532560,152.92059180
4838,"Oro Province",171,PG,"Papua New Guinea",NPP,,-8.89880630,148.18929210
4837,"Port Moresby",171,PG,"Papua New Guinea",NCD,,-9.44380040,147.18026710
4836,"Sandaun Province",171,PG,"Papua New Guinea",SAN,,-3.71261790,141.68342750
4844,"Southern Highlands Province",171,PG,"Papua New Guinea",SHM,,-6.41790830,143.56356370
4830,"West New Britain Province",171,PG,"Papua New Guinea",WBK,,-5.70474320,150.02594660
4840,"Western Highlands Province",171,PG,"Papua New Guinea",WHM,,-5.62681280,144.25931180
4850,"Western Province",171,PG,"Papua New Guinea",WPD,,,
2785,"Alto Paraguay Department",172,PY,Paraguay,16,,-20.08525080,-59.47209040
2784,"Alto Paraná Department",172,PY,Paraguay,10,,-25.60755460,-54.96118360
2782,"Amambay Department",172,PY,Paraguay,13,,-22.55902720,-56.02499820
5221,Asuncion,172,PY,Paraguay,ASU,capital,-25.29682970,-57.68066230
2780,"Boquerón Department",172,PY,Paraguay,19,,-21.74492540,-60.95400730
2773,Caaguazú,172,PY,Paraguay,5,,-25.46458180,-56.01385100
2775,Caazapá,172,PY,Paraguay,6,,-26.18277130,-56.37123270
2771,Canindeyú,172,PY,Paraguay,14,,-24.13787350,-55.66896360
2777,"Central Department",172,PY,Paraguay,11,,36.15592290,-95.96620750
2779,"Concepción Department",172,PY,Paraguay,1,,-23.42142640,-57.43444510
2783,"Cordillera Department",172,PY,Paraguay,3,,-25.22894910,-57.01116810
2772,"Guairá Department",172,PY,Paraguay,4,,-25.88109320,-56.29293810
2778,Itapúa,172,PY,Paraguay,7,,-26.79236230,-55.66896360
2786,"Misiones Department",172,PY,Paraguay,8,,-26.84335120,-57.10131880
2781,"Ñeembucú Department",172,PY,Paraguay,12,,-27.02991140,-57.82539500
2774,"Paraguarí Department",172,PY,Paraguay,9,,-25.62621740,-57.15206420
2770,"Presidente Hayes Department",172,PY,Paraguay,15,,-23.35126050,-58.73736340
2776,"San Pedro Department",172,PY,Paraguay,2,,-24.19486680,-56.56164700
3685,Amazonas,173,PE,Peru,AMA,,,
3680,Áncash,173,PE,Peru,ANC,,-9.32504970,-77.56194190
3699,Apurímac,173,PE,Peru,APU,,-14.05045330,-73.08774900
3681,Arequipa,173,PE,Peru,ARE,,-16.40904740,-71.53745100
3692,Ayacucho,173,PE,Peru,AYA,,-13.16387370,-74.22356410
3688,Cajamarca,173,PE,Peru,CAJ,,-7.16174650,-78.51278550
3701,Callao,173,PE,Peru,CAL,,-12.05084910,-77.12598430
3691,Cusco,173,PE,Peru,CUS,,-13.53195000,-71.96746260
3679,Huancavelica,173,PE,Peru,HUV,,-12.78619780,-74.97640240
3687,Huanuco,173,PE,Peru,HUC,,-9.92076480,-76.24108430
3700,Ica,173,PE,Peru,ICA,,42.35288320,-71.04300970
3693,Junín,173,PE,Peru,JUN,,-11.15819250,-75.99263060
3683,"La Libertad",173,PE,Peru,LAL,,13.49069700,-89.30846070
3702,Lambayeque,173,PE,Peru,LAM,,-6.71976660,-79.90807570
3695,Lima,173,PE,Peru,LIM,,-12.04637310,-77.04275400
4922,Loreto,173,PE,Peru,LOR,,-4.37416430,-76.13042640
3678,"Madre de Dios",173,PE,Peru,MDD,,-11.76687050,-70.81199530
3698,Moquegua,173,PE,Peru,MOQ,,-17.19273610,-70.93281380
3686,Pasco,173,PE,Peru,PAS,,46.23050490,-119.09223160
3697,Piura,173,PE,Peru,PIU,,-5.17828840,-80.65488820
3682,Puno,173,PE,Peru,PUN,,-15.84022180,-70.02188050
3694,"San Martín",173,PE,Peru,SAM,,37.08494640,-121.61022160
3696,Tacna,173,PE,Peru,TAC,,-18.00656790,-70.24627410
3689,Tumbes,173,PE,Peru,TUM,,-3.55649210,-80.42708850
3684,Ucayali,173,PE,Peru,UCA,,-9.82511830,-73.08774900
1324,Abra,174,PH,Philippines,ABR,province,42.49708300,-96.38441000
1323,"Agusan del Norte",174,PH,Philippines,AGN,province,8.94562590,125.53192340
1326,"Agusan del Sur",174,PH,Philippines,AGS,province,8.04638880,126.06153840
1331,Aklan,174,PH,Philippines,AKL,province,11.81661090,122.09415410
1337,Albay,174,PH,Philippines,ALB,province,13.17748270,123.52800720
1336,Antique,174,PH,Philippines,ANT,province,37.03586950,-95.63616940
1334,Apayao,174,PH,Philippines,APA,province,18.01203040,121.17103890
1341,Aurora,174,PH,Philippines,AUR,province,36.97089100,-93.71797900
1316,"Autonomous Region in Muslim Mindanao",174,PH,Philippines,14,region,6.95683130,124.24215970
1346,Basilan,174,PH,Philippines,BAS,province,6.42963490,121.98701650
1344,Bataan,174,PH,Philippines,BAN,province,14.64168420,120.48184460
1352,Batanes,174,PH,Philippines,BTN,province,20.44850740,121.97081290
1359,Batangas,174,PH,Philippines,BTG,province,13.75646510,121.05830760
1363,Benguet,174,PH,Philippines,BEN,province,16.55772570,120.80394740
1304,Bicol,174,PH,Philippines,05,region,13.42098850,123.41367360
1274,Biliran,174,PH,Philippines,BIL,province,11.58331520,124.46418480
1272,Bohol,174,PH,Philippines,BOH,province,9.84999110,124.14354270
1270,Bukidnon,174,PH,Philippines,BUK,province,8.05150540,124.92299460
1278,Bulacan,174,PH,Philippines,BUL,province,14.79427350,120.87990080
1279,Cagayan,174,PH,Philippines,CAG,province,18.24896290,121.87878330
1342,"Cagayan Valley",174,PH,Philippines,02,region,16.97537580,121.81070790
1294,Calabarzon,174,PH,Philippines,40,region,14.10078030,121.07937050
1283,"Camarines Norte",174,PH,Philippines,CAN,province,14.13902650,122.76330360
1287,"Camarines Sur",174,PH,Philippines,CAS,province,13.52501970,123.34861470
1285,Camiguin,174,PH,Philippines,CAM,province,9.17321640,124.72987650
1292,Capiz,174,PH,Philippines,CAP,province,11.55288160,122.74072300
1314,Caraga,174,PH,Philippines,13,region,8.80145620,125.74068820
1301,Catanduanes,174,PH,Philippines,CAT,province,13.70886840,124.24215970
1307,Cavite,174,PH,Philippines,CAV,province,14.47912970,120.89696340
1306,Cebu,174,PH,Philippines,CEB,province,10.31569920,123.88543660
1345,"Central Luzon",174,PH,Philippines,03,region,15.48277220,120.71200230
1308,"Central Visayas",174,PH,Philippines,07,region,9.81687500,124.06414190
1311,"Compostela Valley",174,PH,Philippines,COM,province,7.51251500,126.17626150
1335,"Cordillera Administrative",174,PH,Philippines,15,region,17.35125420,121.17188510
1320,Cotabato,174,PH,Philippines,NCO,province,7.20466680,124.23104390
1340,Davao,174,PH,Philippines,11,region,7.30416220,126.08934060
1319,"Davao del Norte",174,PH,Philippines,DAV,province,7.56176990,125.65328480
1318,"Davao del Sur",174,PH,Philippines,DAS,province,6.76626870,125.32842690
1309,"Davao Occidental",174,PH,Philippines,DVO,province,6.09413960,125.60954740
1289,"Davao Oriental",174,PH,Philippines,DAO,province,7.31715850,126.54198870
1291,"Dinagat Islands",174,PH,Philippines,DIN,province,10.12818160,125.60954740
1290,"Eastern Samar",174,PH,Philippines,EAS,province,11.50007310,125.49999080
1322,"Eastern Visayas",174,PH,Philippines,08,region,12.24455330,125.03881640
1303,Guimaras,174,PH,Philippines,GUI,province,10.59286610,122.63250810
1300,Ifugao,174,PH,Philippines,IFU,province,16.83307920,121.17103890
1355,Ilocos,174,PH,Philippines,01,region,16.08321440,120.61998950
1298,"Ilocos Norte",174,PH,Philippines,ILN,province,18.16472810,120.71155920
1321,"Ilocos Sur",174,PH,Philippines,ILS,province,17.22786640,120.57395790
1315,Iloilo,174,PH,Philippines,ILI,province,10.72015010,122.56210630
1313,Isabela,174,PH,Philippines,ISA,province,18.50077590,-67.02434620
1312,Kalinga,174,PH,Philippines,KAL,province,17.47404220,121.35416310
1317,"La Union",174,PH,Philippines,LUN,province,38.87668780,-77.12809120
1328,Laguna,174,PH,Philippines,LAG,province,33.54271890,-117.78535680
1327,"Lanao del Norte",174,PH,Philippines,LAN,province,7.87218110,123.88577470
1333,"Lanao del Sur",174,PH,Philippines,LAS,province,7.82317600,124.41982430
1332,Leyte,174,PH,Philippines,LEY,province,10.86245360,124.88111950
1330,Maguindanao,174,PH,Philippines,MAG,province,6.94225810,124.41982430
1329,Marinduque,174,PH,Philippines,MAD,province,13.47671710,121.90321920
1338,Masbate,174,PH,Philippines,MAS,province,12.35743460,123.55040760
1347,"Metro Manila",174,PH,Philippines,NCR,province,14.60905370,121.02225650
1299,Mimaropa,174,PH,Philippines,41,region,9.84320650,118.73647830
1343,"Misamis Occidental",174,PH,Philippines,MSC,province,8.33749030,123.70706190
1348,"Misamis Oriental",174,PH,Philippines,MSR,province,8.50455580,124.62195920
1353,"Mountain Province",174,PH,Philippines,MOU,province,40.70754370,-73.95010330
1351,"Negros Occidental",174,PH,Philippines,NEC,province,10.29256090,123.02465180
1350,"Negros Oriental",174,PH,Philippines,NER,province,9.62820830,122.98883190
1339,"Northern Mindanao",174,PH,Philippines,10,region,8.02016350,124.68565090
1349,"Northern Samar",174,PH,Philippines,NSA,province,12.36131990,124.77407930
1360,"Nueva Ecija",174,PH,Philippines,NUE,province,15.57837500,121.11126150
1358,"Nueva Vizcaya",174,PH,Philippines,NUV,province,16.33011070,121.17103890
1356,"Occidental Mindoro",174,PH,Philippines,MDC,province,13.10241110,120.76512840
1354,"Oriental Mindoro",174,PH,Philippines,MDR,province,13.05645980,121.40694170
1361,Palawan,174,PH,Philippines,PLW,province,9.83494930,118.73836150
1365,Pampanga,174,PH,Philippines,PAM,province,15.07940900,120.61998950
1364,Pangasinan,174,PH,Philippines,PAN,province,15.89490550,120.28631830
1275,Quezon,174,PH,Philippines,QUE,province,14.03139060,122.11309090
1273,Quirino,174,PH,Philippines,QUI,province,16.27004240,121.53700030
1271,Rizal,174,PH,Philippines,RIZ,province,14.60374460,121.30840880
1269,Romblon,174,PH,Philippines,ROM,province,12.57780160,122.26914600
1277,Sarangani,174,PH,Philippines,SAR,province,5.92671750,124.99475100
1276,Siquijor,174,PH,Philippines,SIG,province,9.19987790,123.59519250
1310,Soccsksargen,174,PH,Philippines,12,region,6.27069180,124.68565090
1281,Sorsogon,174,PH,Philippines,SOR,province,12.99270950,124.01474640
1280,"South Cotabato",174,PH,Philippines,SCO,province,6.33575650,124.77407930
1284,"Southern Leyte",174,PH,Philippines,SLE,province,10.33462060,125.17087410
1282,"Sultan Kudarat",174,PH,Philippines,SUK,province,6.50694010,124.41982430
1288,Sulu,174,PH,Philippines,SLU,province,5.97490110,121.03351000
1286,"Surigao del Norte",174,PH,Philippines,SUN,province,9.51482800,125.69699840
1296,"Surigao del Sur",174,PH,Philippines,SUR,province,8.54049060,126.11447580
1295,Tarlac,174,PH,Philippines,TAR,province,15.47547860,120.59634920
1293,Tawi-Tawi,174,PH,Philippines,TAW,province,5.13381100,119.95092600
5115,"Western Samar",174,PH,Philippines,WSA,province,12.00002060,124.99124520
1305,"Western Visayas",174,PH,Philippines,06,region,11.00498360,122.53727410
1297,Zambales,174,PH,Philippines,ZMB,province,15.50817660,119.96978080
1302,"Zamboanga del Norte",174,PH,Philippines,ZAN,province,8.38862820,123.16888830
1357,"Zamboanga del Sur",174,PH,Philippines,ZAS,province,7.83830540,123.29666570
1325,"Zamboanga Peninsula",174,PH,Philippines,09,region,8.15407700,123.25879300
1362,"Zamboanga Sibugay",174,PH,Philippines,ZSI,province,7.52252470,122.31075170
1634,"Greater Poland Voivodeship",176,PL,Poland,WP,,52.27998600,17.35229390
1625,"Kuyavian-Pomeranian Voivodeship",176,PL,Poland,KP,,53.16483630,18.48342240
1635,"Lesser Poland Voivodeship",176,PL,Poland,MA,,49.72253060,20.25033580
1629,"Lower Silesian Voivodeship",176,PL,Poland,DS,,51.13398610,16.88419610
1638,"Lublin Voivodeship",176,PL,Poland,LU,,51.24935190,23.10113920
1631,"Lubusz Voivodeship",176,PL,Poland,LB,,52.22746120,15.25591030
1636,"Łódź Voivodeship",176,PL,Poland,LD,,51.46347710,19.17269740
1637,"Masovian Voivodeship",176,PL,Poland,MZ,,51.89271820,21.00216790
1622,"Opole Voivodeship",176,PL,Poland,OP,,50.80037610,17.93798900
1626,"Podkarpackie Voivodeship",176,PL,Poland,PK,,50.05747490,22.08956910
1632,"Podlaskie Voivodeship",176,PL,Poland,PD,,53.06971590,22.96746390
1624,"Pomeranian Voivodeship",176,PL,Poland,PM,,54.29442520,18.15311640
1623,"Silesian Voivodeship",176,PL,Poland,SL,,50.57165950,19.32197680
1630,"Świętokrzyskie Voivodeship",176,PL,Poland,SK,,50.62610410,20.94062790
1628,"Warmian-Masurian Voivodeship",176,PL,Poland,WN,,53.86711170,20.70278610
1633,"West Pomeranian Voivodeship",176,PL,Poland,ZP,,53.46578910,15.18225810
2233,Açores,177,PT,Portugal,20,,37.74124880,-25.67559440
2235,Aveiro,177,PT,Portugal,01,,40.72090230,-8.57210160
2230,Beja,177,PT,Portugal,02,,37.96877860,-7.87216000
2244,Braga,177,PT,Portugal,03,,41.55038800,-8.42613010
2229,Bragança,177,PT,Portugal,04,,41.80616520,-6.75674270
2241,"Castelo Branco",177,PT,Portugal,05,,39.86313230,-7.48141630
2246,Coimbra,177,PT,Portugal,06,,40.20579940,-8.41369000
2236,Évora,177,PT,Portugal,07,,38.57444680,-7.90765530
2239,Faro,177,PT,Portugal,08,,37.01935480,-7.93043970
4859,Guarda,177,PT,Portugal,09,,40.53859720,-7.26757720
2240,Leiria,177,PT,Portugal,10,,39.77095320,-8.79218360
2228,Lisbon,177,PT,Portugal,11,,38.72232630,-9.13927140
2231,Madeira,177,PT,Portugal,30,,32.76070740,-16.95947230
2232,Portalegre,177,PT,Portugal,12,,39.29670860,-7.42847550
2243,Porto,177,PT,Portugal,13,,41.14766290,-8.60789730
2238,Santarém,177,PT,Portugal,14,,39.23666870,-8.68599440
2242,Setúbal,177,PT,Portugal,15,,38.52409330,-8.89258760
2245,"Viana do Castelo",177,PT,Portugal,16,,41.69180460,-8.83445100
2234,"Vila Real",177,PT,Portugal,17,,41.30035270,-7.74572740
2237,Viseu,177,PT,Portugal,18,,40.65884240,-7.91475600
5134,Adjuntas,178,PR,"Puerto Rico",001,municipalities,18.16348480,-66.72315800
5135,Aguada,178,PR,"Puerto Rico",003,municipalities,18.38015790,-67.18870400
5136,Aguadilla,178,PR,"Puerto Rico",005,municipalities,18.42744540,-67.15406980
5137,"Aguas Buenas",178,PR,"Puerto Rico",007,municipalities,18.25689890,-66.10294420
5138,Aibonito,178,PR,"Puerto Rico",009,municipalities,18.13995940,-66.26600160
5139,Añasco,178,PR,"Puerto Rico",011,municipalities,18.28544760,-67.14029350
5140,Arecibo,178,PR,"Puerto Rico",013,municipalities,18.47051370,-66.72184720
5081,Arecibo,178,PR,"Puerto Rico",AR,region,18.47055556,-66.72083333
5141,Arroyo,178,PR,"Puerto Rico",015,municipalities,17.99642200,-66.09248790
5142,Barceloneta,178,PR,"Puerto Rico",017,municipalities,41.38010610,2.18969570
5143,Barranquitas,178,PR,"Puerto Rico",019,municipalities,18.18662420,-66.30628020
5076,Bayamon,178,PR,"Puerto Rico",BY,region,18.17777778,-66.11333333
5144,Bayamón,178,PR,"Puerto Rico",021,municipalities,18.38939600,-66.16532240
5145,"Cabo Rojo",178,PR,"Puerto Rico",023,municipalities,18.08662650,-67.14573470
5079,Caguas,178,PR,"Puerto Rico",CG,region,18.23333333,-66.03333333
5146,Caguas,178,PR,"Puerto Rico",025,municipalities,18.23879950,-66.03524900
5147,Camuy,178,PR,"Puerto Rico",027,municipalities,18.48383300,-66.84489940
5148,Canóvanas,178,PR,"Puerto Rico",029,municipalities,18.37487480,-65.89975330
5077,Carolina,178,PR,"Puerto Rico",CL,region,18.38888889,-65.96666667
5149,Carolina,178,PR,"Puerto Rico",031,municipalities,18.36808770,-66.04247340
5150,Cataño,178,PR,"Puerto Rico",033,municipalities,18.44653550,-66.13557750
5151,Cayey,178,PR,"Puerto Rico",035,municipalities,18.11190510,-66.16600000
5152,Ceiba,178,PR,"Puerto Rico",037,municipalities,18.24751770,-65.90849530
5153,Ciales,178,PR,"Puerto Rico",039,municipalities,18.33606220,-66.46878230
5154,Cidra,178,PR,"Puerto Rico",041,municipalities,18.17579140,-66.16127790
5155,Coamo,178,PR,"Puerto Rico",043,municipalities,18.07996160,-66.35794730
5156,Comerío,178,PR,"Puerto Rico",045,municipalities,18.21920010,-66.22560220
5157,Corozal,178,PR,"Puerto Rico",047,municipalities,18.40308020,-88.39675360
5158,Culebra,178,PR,"Puerto Rico",049,municipalities,18.31039400,-65.30307050
5159,Dorado,178,PR,"Puerto Rico",051,municipalities,43.14805560,-77.57722220
5160,Fajardo,178,PR,"Puerto Rico",053,municipalities,18.32521480,-65.65393560
5161,Florida,178,PR,"Puerto Rico",054,municipalities,27.66482740,-81.51575350
5162,Guánica,178,PR,"Puerto Rico",055,municipalities,17.97251450,-66.90862640
5163,Guayama,178,PR,"Puerto Rico",057,municipalities,17.98413280,-66.11377670
5164,Guayanilla,178,PR,"Puerto Rico",059,municipalities,18.01913140,-66.79184200
5080,Guaynabo,178,PR,"Puerto Rico",GN,region,18.36666667,-66.10000000
5165,Guaynabo,178,PR,"Puerto Rico",061,municipalities,18.36129540,-66.11029570
5166,Gurabo,178,PR,"Puerto Rico",063,municipalities,18.25439870,-65.97294210
5167,Hatillo,178,PR,"Puerto Rico",065,municipalities,18.42846420,-66.78753210
5168,Hormigueros,178,PR,"Puerto Rico",067,municipalities,18.13346380,-67.11281230
5169,Humacao,178,PR,"Puerto Rico",069,municipalities,18.15157360,-65.82485290
5170,Isabela,178,PR,"Puerto Rico",071,municipalities,16.97537580,121.81070790
5171,Jayuya,178,PR,"Puerto Rico",073,municipalities,18.21856740,-66.59156170
5172,"Juana Díaz",178,PR,"Puerto Rico",075,municipalities,18.05343720,-66.50750790
5173,Juncos,178,PR,"Puerto Rico",077,municipalities,18.22745580,-65.92099700
5174,Lajas,178,PR,"Puerto Rico",079,municipalities,18.04996200,-67.05934490
5175,Lares,178,PR,"Puerto Rico",081,municipalities,34.02508020,-118.45945930
5176,"Las Marías",178,PR,"Puerto Rico",083,municipalities,35.83822380,-78.61035660
5177,"Las Piedras",178,PR,"Puerto Rico",085,municipalities,18.18557530,-65.87362450
5178,Loíza,178,PR,"Puerto Rico",087,municipalities,18.43299040,-65.87836000
5179,Luquillo,178,PR,"Puerto Rico",089,municipalities,18.37245070,-65.71655110
5180,Manatí,178,PR,"Puerto Rico",091,municipalities,18.41812150,-66.52627830
5181,Maricao,178,PR,"Puerto Rico",093,municipalities,18.18079020,-66.97990010
5182,Maunabo,178,PR,"Puerto Rico",095,municipalities,18.00718850,-65.89932890
5083,Mayagüez,178,PR,"Puerto Rico",MG,region,18.20111111,-67.13972222
5183,Mayagüez,178,PR,"Puerto Rico",097,municipalities,18.20134520,-67.14515490
5184,Moca,178,PR,"Puerto Rico",099,municipalities,18.39679290,-67.14790350
5185,Morovis,178,PR,"Puerto Rico",101,municipalities,18.32578500,-66.40655920
5186,Naguabo,178,PR,"Puerto Rico",103,municipalities,18.21162470,-65.73488410
5187,Naranjito,178,PR,"Puerto Rico",105,municipalities,18.30078610,-66.24489040
5188,Orocovis,178,PR,"Puerto Rico",107,municipalities,18.22692240,-66.39116860
5189,Patillas,178,PR,"Puerto Rico",109,municipalities,18.00373810,-66.01340590
5190,Peñuelas,178,PR,"Puerto Rico",111,municipalities,18.06335770,-66.72738960
5078,Ponce,178,PR,"Puerto Rico",PO,region,18.00000000,-66.61666667
5191,Ponce,178,PR,"Puerto Rico",113,municipalities,18.01107680,-66.61406160
5192,Quebradillas,178,PR,"Puerto Rico",115,municipalities,18.47383300,-66.93851200
5193,Rincón,178,PR,"Puerto Rico",117,municipalities,18.34015140,-67.24994590
5194,"Río Grande",178,PR,"Puerto Rico",119,municipalities,28.81063826,-101.83538780
5195,"Sabana Grande",178,PR,"Puerto Rico",121,municipalities,18.07773920,-66.96045490
5196,Salinas,178,PR,"Puerto Rico",123,municipalities,36.67773720,-121.65550130
5197,"San Germán",178,PR,"Puerto Rico",125,municipalities,18.08070820,-67.04110960
5198,"San Juan",178,PR,"Puerto Rico",127,municipalities,18.46320300,-66.11475710
5075,"San Juan",178,PR,"Puerto Rico",SJ,region,18.45000000,-66.06666667
5199,"San Lorenzo",178,PR,"Puerto Rico",129,municipalities,18.18869120,-65.97658620
5200,"San Sebastián",178,PR,"Puerto Rico",131,municipalities,43.31833400,-1.98123130
5201,"Santa Isabel",178,PR,"Puerto Rico",133,municipalities,17.96607750,-66.40489200
5202,"Toa Alta",178,PR,"Puerto Rico",135,municipalities,18.38828230,-66.24822370
5082,"Toa Baja",178,PR,"Puerto Rico",TB,region,18.44388900,-66.25972200
5203,"Toa Baja",178,PR,"Puerto Rico",137,municipalities,18.44447090,-66.25432930
5084,"Trujillo Alto",178,PR,"Puerto Rico",TA,region,18.36277800,-66.01750000
5204,"Trujillo Alto",178,PR,"Puerto Rico",139,municipalities,18.35467190,-66.00738760
5205,Utuado,178,PR,"Puerto Rico",141,municipalities,18.26550950,-66.70045190
5206,"Vega Alta",178,PR,"Puerto Rico",143,municipalities,18.41217030,-66.33128050
5207,"Vega Baja",178,PR,"Puerto Rico",145,municipalities,18.44614590,-66.40419670
5208,Vieques,178,PR,"Puerto Rico",147,municipalities,18.12628540,-65.44009850
5209,Villalba,178,PR,"Puerto Rico",149,municipalities,18.12175540,-66.49857870
5210,Yabucoa,178,PR,"Puerto Rico",151,municipalities,18.05052010,-65.87932880
5211,Yauco,178,PR,"Puerto Rico",153,municipalities,18.03496400,-66.84989830
3182,"Al Daayen",179,QA,Qatar,ZA,,25.57845590,51.48213870
3183,"Al Khor",179,QA,Qatar,KH,,25.68040780,51.49685020
3177,"Al Rayyan Municipality",179,QA,Qatar,RA,,25.25225510,51.43887130
3179,"Al Wakrah",179,QA,Qatar,WA,,25.16593140,51.59755240
3178,Al-Shahaniya,179,QA,Qatar,SH,,25.41063860,51.18460250
3181,Doha,179,QA,Qatar,DA,,25.28544730,51.53103980
3180,"Madinat ash Shamal",179,QA,Qatar,MS,,26.11827430,51.21572650
3184,"Umm Salal Municipality",179,QA,Qatar,US,,25.48752420,51.39656800
4724,Alba,181,RO,Romania,AB,,44.70091530,8.03569110
4739,"Arad County",181,RO,Romania,AR,,46.22836510,21.65978190
4722,Arges,181,RO,Romania,AG,,45.07225270,24.81427260
4744,"Bacău County",181,RO,Romania,BC,,46.32581840,26.66237800
4723,"Bihor County",181,RO,Romania,BH,,47.01575160,22.17226600
4733,"Bistrița-Năsăud County",181,RO,Romania,BN,,47.24861070,24.53228140
4740,"Botoșani County",181,RO,Romania,BT,,47.89240420,26.75917810
4736,Braila,181,RO,Romania,BR,,45.26524630,27.95947140
4759,"Brașov County",181,RO,Romania,BV,,45.77818440,25.22258000
4730,Bucharest,181,RO,Romania,B,,44.42676740,26.10253840
4756,"Buzău County",181,RO,Romania,BZ,,45.33509120,26.71075780
4732,"Călărași County",181,RO,Romania,CL,,44.36587150,26.75486070
4753,"Caraș-Severin County",181,RO,Romania,CS,,45.11396460,22.07409930
4734,"Cluj County",181,RO,Romania,CJ,,46.79417970,23.61214920
4737,"Constanța County",181,RO,Romania,CT,,44.21298700,28.25500550
4754,"Covasna County",181,RO,Romania,CV,,45.94263470,25.89189840
4745,"Dâmbovița County",181,RO,Romania,DB,,44.92898930,25.42538500
4742,"Dolj County",181,RO,Romania,DJ,,44.16230220,23.63250540
4747,"Galați County",181,RO,Romania,GL,,45.78005690,27.82515760
4726,"Giurgiu County",181,RO,Romania,GR,,43.90370760,25.96992650
4750,"Gorj County",181,RO,Romania,GJ,,44.94855950,23.24270790
4749,"Harghita County",181,RO,Romania,HR,,46.49285070,25.64566960
4721,"Hunedoara County",181,RO,Romania,HD,,45.79363800,22.99759930
4743,"Ialomița County",181,RO,Romania,IL,,44.60313300,27.37899140
4735,"Iași County",181,RO,Romania,IS,,47.26796530,27.21856620
4725,"Ilfov County",181,RO,Romania,IF,,44.53554800,26.23248860
4760,"Maramureș County",181,RO,Romania,MM,,46.55699040,24.67232150
4751,"Mehedinți County",181,RO,Romania,MH,,44.55150530,22.90441570
4915,"Mureș County",181,RO,Romania,MS,,46.55699040,24.67232150
4731,"Neamț County",181,RO,Romania,NT,,46.97586850,26.38187640
4738,"Olt County",181,RO,Romania,OT,,44.20079700,24.50229810
4729,"Prahova County",181,RO,Romania,PH,,45.08919060,26.08293130
4741,"Sălaj County",181,RO,Romania,SJ,,47.20908130,23.21219010
4746,"Satu Mare County",181,RO,Romania,SM,,47.76689050,22.92413770
4755,"Sibiu County",181,RO,Romania,SB,,45.92691060,24.22548070
4720,"Suceava County",181,RO,Romania,SV,,47.55055480,25.74106200
4728,"Teleorman County",181,RO,Romania,TR,,44.01604910,25.29866280
4748,"Timiș County",181,RO,Romania,TM,,45.81389020,21.33310550
4727,"Tulcea County",181,RO,Romania,TL,,45.04505650,29.03249120
4757,"Vâlcea County",181,RO,Romania,VL,,45.07980510,24.08352830
4752,"Vaslui County",181,RO,Romania,VS,,46.46310590,27.73980310
4758,"Vrancea County",181,RO,Romania,VN,,45.81348760,27.06575310
1911,"Altai Krai",182,RU,Russia,ALT,,51.79362980,82.67585960
1876,"Altai Republic",182,RU,Russia,AL,,50.61819240,86.21993080
1858,"Amur Oblast",182,RU,Russia,AMU,,54.60350650,127.48017210
1849,Arkhangelsk,182,RU,Russia,ARK,,64.54585490,40.55057690
1866,"Astrakhan Oblast",182,RU,Russia,AST,,46.13211660,48.06101150
1903,"Belgorod Oblast",182,RU,Russia,BEL,,50.71069260,37.75333770
1867,"Bryansk Oblast",182,RU,Russia,BRY,,53.04085990,33.26909000
1893,"Chechen Republic",182,RU,Russia,CE,,43.40233010,45.71874680
1845,"Chelyabinsk Oblast",182,RU,Russia,CHE,,54.43194220,60.87889630
1859,"Chukotka Autonomous Okrug",182,RU,Russia,CHU,,65.62983550,171.69521590
1914,"Chuvash Republic",182,RU,Russia,CU,,55.55959920,46.92835350
1880,Irkutsk,182,RU,Russia,IRK,,52.28548340,104.28902220
1864,"Ivanovo Oblast",182,RU,Russia,IVA,,57.10568540,41.48300840
1835,"Jewish Autonomous Oblast",182,RU,Russia,YEV,,48.48081470,131.76573670
1892,"Kabardino-Balkar Republic",182,RU,Russia,KB,,43.39324690,43.56284980
1902,Kaliningrad,182,RU,Russia,KGD,,54.71042640,20.45221440
1844,"Kaluga Oblast",182,RU,Russia,KLU,,54.38726660,35.18890940
1865,"Kamchatka Krai",182,RU,Russia,KAM,,61.43439810,166.78841310
1869,"Karachay-Cherkess Republic",182,RU,Russia,KC,,43.88451430,41.73039390
1897,"Kemerovo Oblast",182,RU,Russia,KEM,,54.75746480,87.40552880
1873,"Khabarovsk Krai",182,RU,Russia,KHA,,50.58884310,135.00000000
1838,"Khanty-Mansi Autonomous Okrug",182,RU,Russia,KHM,,62.22870620,70.64100570
1890,"Kirov Oblast",182,RU,Russia,KIR,,58.41985290,50.20972480
1899,"Komi Republic",182,RU,Russia,KO,,63.86305390,54.83126900
1910,"Kostroma Oblast",182,RU,Russia,KOS,,58.55010690,43.95411020
1891,"Krasnodar Krai",182,RU,Russia,KDA,,45.64152890,39.70559770
1840,"Krasnoyarsk Krai",182,RU,Russia,KYA,,64.24797580,95.11041760
1915,"Kurgan Oblast",182,RU,Russia,KGN,,55.44815480,65.11809750
1855,"Kursk Oblast",182,RU,Russia,KRS,,51.76340260,35.38118120
1896,"Leningrad Oblast",182,RU,Russia,LEN,,60.07932080,31.89266450
1889,"Lipetsk Oblast",182,RU,Russia,LIP,,52.52647020,39.20322690
1839,"Magadan Oblast",182,RU,Russia,MAG,,62.66434170,153.91499100
1870,"Mari El Republic",182,RU,Russia,ME,,56.43845700,47.96077580
1901,Moscow,182,RU,Russia,MOW,,55.75582600,37.61729990
1882,"Moscow Oblast",182,RU,Russia,MOS,,55.34039600,38.29176510
1843,"Murmansk Oblast",182,RU,Russia,MUR,,67.84426740,35.08841020
1836,"Nenets Autonomous Okrug",182,RU,Russia,NEN,,67.60783370,57.63383310
1857,"Nizhny Novgorod Oblast",182,RU,Russia,NIZ,,55.79951590,44.02967690
1834,"Novgorod Oblast",182,RU,Russia,NGR,,58.24275520,32.56651900
1888,Novosibirsk,182,RU,Russia,NVS,,54.98326930,82.89638310
1846,"Omsk Oblast",182,RU,Russia,OMS,,55.05546690,73.31673420
1886,"Orenburg Oblast",182,RU,Russia,ORE,,51.76340260,54.61881880
1908,"Oryol Oblast",182,RU,Russia,ORL,,52.78564140,36.92423440
1909,"Penza Oblast",182,RU,Russia,PNZ,,53.14121050,44.09400480
1871,"Perm Krai",182,RU,Russia,PER,,58.82319290,56.58724810
1833,"Primorsky Krai",182,RU,Russia,PRI,,45.05256410,135.00000000
1863,"Pskov Oblast",182,RU,Russia,PSK,,56.77085990,29.09400900
1852,"Republic of Adygea",182,RU,Russia,AD,,44.82291550,40.17544630
1854,"Republic of Bashkortostan",182,RU,Russia,BA,,54.23121720,56.16452570
1842,"Republic of Buryatia",182,RU,Russia,BU,,54.83311460,112.40605300
1850,"Republic of Dagestan",182,RU,Russia,DA,,42.14318860,47.09497990
1884,"Republic of Ingushetia",182,RU,Russia,IN,,43.40516980,44.82029990
1883,"Republic of Kalmykia",182,RU,Russia,KL,,46.18671760,45.00000000
1841,"Republic of Karelia",182,RU,Russia,KR,,63.15587020,32.99055520
1877,"Republic of Khakassia",182,RU,Russia,KK,,53.04522810,90.39821450
1898,"Republic of Mordovia",182,RU,Russia,MO,,54.23694410,44.06839700
1853,"Republic of North Ossetia-Alania",182,RU,Russia,SE,,43.04513020,44.28709720
1861,"Republic of Tatarstan",182,RU,Russia,TA,,55.18023640,50.72639450
1837,"Rostov Oblast",182,RU,Russia,ROS,,47.68532470,41.82589520
1905,"Ryazan Oblast",182,RU,Russia,RYA,,54.38759640,41.25956610
1879,"Saint Petersburg",182,RU,Russia,SPE,,59.93105840,30.36090960
1848,"Sakha Republic",182,RU,Russia,SA,,66.76134510,124.12375300
1875,Sakhalin,182,RU,Russia,SAK,,50.69098480,142.95056890
1862,"Samara Oblast",182,RU,Russia,SAM,,53.41838390,50.47255280
1887,"Saratov Oblast",182,RU,Russia,SAR,,51.83692630,46.75393970
1885,"Smolensk Oblast",182,RU,Russia,SMO,,54.98829940,32.66773780
1868,"Stavropol Krai",182,RU,Russia,STA,,44.66809930,43.52021400
1894,Sverdlovsk,182,RU,Russia,SVE,,56.84309930,60.64540860
1878,"Tambov Oblast",182,RU,Russia,TAM,,52.64165890,41.42164510
1872,"Tomsk Oblast",182,RU,Russia,TOM,,58.89698820,82.67655000
1895,"Tula Oblast",182,RU,Russia,TUL,,54.16376800,37.56495070
1900,"Tuva Republic",182,RU,Russia,TY,,51.88726690,95.62601720
1860,"Tver Oblast",182,RU,Russia,TVE,,57.00216540,33.98531420
1907,"Tyumen Oblast",182,RU,Russia,TYU,,56.96343870,66.94827800
1913,"Udmurt Republic",182,RU,Russia,UD,,57.06702180,53.02779480
1856,"Ulyanovsk Oblast",182,RU,Russia,ULY,,53.97933570,47.77624250
1881,"Vladimir Oblast",182,RU,Russia,VLA,,56.15534650,40.59266850
4916,"Volgograd Oblast",182,RU,Russia,VGG,,49.25873930,39.81544630
1874,"Vologda Oblast",182,RU,Russia,VLG,,59.87067110,40.65554110
1906,"Voronezh Oblast",182,RU,Russia,VOR,,50.85897130,39.86443740
1847,"Yamalo-Nenets Autonomous Okrug",182,RU,Russia,YAN,,66.06530570,76.93451930
1851,"Yaroslavl Oblast",182,RU,Russia,YAR,,57.89915230,38.83886330
1904,"Zabaykalsky Krai",182,RU,Russia,ZAB,,53.09287710,116.96765610
261,"Eastern Province",183,RW,Rwanda,02,,,
262,"Kigali district",183,RW,Rwanda,01,,-1.94407270,30.06188510
263,"Northern Province",183,RW,Rwanda,03,,,
259,"Southern Province",183,RW,Rwanda,05,,,
260,"Western Province",183,RW,Rwanda,04,,,
3833,"Christ Church Nichola Town Parish",185,KN,"Saint Kitts and Nevis",01,,17.36048120,-62.76178370
3832,Nevis,185,KN,"Saint Kitts and Nevis",N,,17.15535580,-62.57960260
3836,"Saint Anne Sandy Point Parish",185,KN,"Saint Kitts and Nevis",02,,17.37253330,-62.84411330
3837,"Saint George Gingerland Parish",185,KN,"Saint Kitts and Nevis",04,,17.12577590,-62.56198110
3835,"Saint James Windward Parish",185,KN,"Saint Kitts and Nevis",05,,17.17696330,-62.57960260
3845,"Saint John Capisterre Parish",185,KN,"Saint Kitts and Nevis",06,,17.38103410,-62.79118330
3840,"Saint John Figtree Parish",185,KN,"Saint Kitts and Nevis",07,,17.11557480,-62.60310040
3841,"Saint Kitts",185,KN,"Saint Kitts and Nevis",K,,17.34337960,-62.75590430
3844,"Saint Mary Cayon Parish",185,KN,"Saint Kitts and Nevis",08,,17.34620710,-62.73826710
3834,"Saint Paul Capisterre Parish",185,KN,"Saint Kitts and Nevis",09,,17.40166830,-62.82573320
3838,"Saint Paul Charlestown Parish",185,KN,"Saint Kitts and Nevis",10,,17.13462970,-62.61338160
3831,"Saint Peter Basseterre Parish",185,KN,"Saint Kitts and Nevis",11,,17.31029110,-62.71475330
3839,"Saint Thomas Lowland Parish",185,KN,"Saint Kitts and Nevis",12,,17.16505130,-62.60897530
3842,"Saint Thomas Middle Island Parish",185,KN,"Saint Kitts and Nevis",13,,17.33488130,-62.80882510
3843,"Trinity Palmetto Point Parish",185,KN,"Saint Kitts and Nevis",15,,17.30635190,-62.76178370
3757,"Anse la Raye Quarter",186,LC,"Saint Lucia",01,,13.94594240,-61.03694680
3761,Canaries,186,LC,"Saint Lucia",12,,28.29156370,-16.62913040
3758,"Castries Quarter",186,LC,"Saint Lucia",02,,14.01010940,-60.98746870
3760,"Choiseul Quarter",186,LC,"Saint Lucia",03,,13.77501540,-61.04859100
3767,"Dauphin Quarter",186,LC,"Saint Lucia",04,,14.01033960,-60.91909880
3756,"Dennery Quarter",186,LC,"Saint Lucia",05,,13.92673930,-60.91909880
3766,"Gros Islet Quarter",186,LC,"Saint Lucia",06,,14.08435780,-60.94527940
3759,"Laborie Quarter",186,LC,"Saint Lucia",07,,13.75227830,-60.99328890
3762,"Micoud Quarter",186,LC,"Saint Lucia",08,,13.82118710,-60.90019340
3765,"Praslin Quarter",186,LC,"Saint Lucia",09,,13.87523920,-60.89946630
3764,"Soufrière Quarter",186,LC,"Saint Lucia",10,,13.85709860,-61.05732480
3763,"Vieux Fort Quarter",186,LC,"Saint Lucia",11,,13.72060800,-60.94964330
3389,"Charlotte Parish",188,VC,"Saint Vincent and the Grenadines",01,,13.21754510,-61.16362440
3388,"Grenadines Parish",188,VC,"Saint Vincent and the Grenadines",06,,13.01229650,-61.22773010
3386,"Saint Andrew Parish",188,VC,"Saint Vincent and the Grenadines",02,,43.02429990,-81.20250000
3387,"Saint David Parish",188,VC,"Saint Vincent and the Grenadines",03,,43.85230950,-79.52366540
3384,"Saint George Parish",188,VC,"Saint Vincent and the Grenadines",04,,42.95760900,-81.32670500
3385,"Saint Patrick Parish",188,VC,"Saint Vincent and the Grenadines",05,,39.75091860,-94.84505560
4763,A'ana,191,WS,Samoa,AA,,-13.89841800,-171.97529950
4761,Aiga-i-le-Tai,191,WS,Samoa,AL,,-13.85137910,-172.03254010
4765,Atua,191,WS,Samoa,AT,,-13.97870530,-171.62542830
4764,Fa'asaleleaga,191,WS,Samoa,FA,,-13.63076380,-172.23659810
4769,Gaga'emauga,191,WS,Samoa,GE,,-13.54286660,-172.36688700
4771,Gaga'ifomauga,191,WS,Samoa,GI,,-13.54680070,-172.49693310
4767,Palauli,191,WS,Samoa,PA,,-13.72945790,-172.45361150
4762,Satupa'itea,191,WS,Samoa,SA,,-13.65382140,-172.61592710
4770,Tuamasaga,191,WS,Samoa,TU,,-13.91635920,-171.82243620
4768,Va'a-o-Fonoti,191,WS,Samoa,VF,,-13.94709030,-171.54318720
4766,Vaisigano,191,WS,Samoa,VS,,-13.54138270,-172.70233830
59,Acquaviva,192,SM,"San Marino",01,,41.86715970,14.74694790
61,"Borgo Maggiore",192,SM,"San Marino",06,,43.95748820,12.45525460
60,Chiesanuova,192,SM,"San Marino",02,,45.42261720,7.65038540
64,Domagnano,192,SM,"San Marino",03,,43.95019290,12.46815370
62,Faetano,192,SM,"San Marino",04,,43.93489670,12.48965540
66,Fiorentino,192,SM,"San Marino",05,,43.90783370,12.45812090
63,Montegiardino,192,SM,"San Marino",08,,43.90529990,12.48105420
58,"San Marino",192,SM,"San Marino",07,,43.94236000,12.45777700
65,Serravalle,192,SM,"San Marino",09,,44.72320840,8.85740050
270,"Príncipe Province",193,ST,"Sao Tome and Principe",P,,1.61393810,7.40569280
271,"São Tomé Province",193,ST,"Sao Tome and Principe",S,,0.33019240,6.73334300
2853,'Asir,194,SA,"Saudi Arabia",14,region,19.09690620,42.86378750
2859,"Al Bahah",194,SA,"Saudi Arabia",11,region,20.27227390,41.44125100
2857,"Al Jawf",194,SA,"Saudi Arabia",12,region,29.88735600,39.32062410
2851,"Al Madinah",194,SA,"Saudi Arabia",03,region,24.84039770,39.32062410
2861,Al-Qassim,194,SA,"Saudi Arabia",05,region,26.20782600,43.48373800
2856,"Eastern Province",194,SA,"Saudi Arabia",04,region,24.04399320,45.65892250
2855,Ha'il,194,SA,"Saudi Arabia",06,region,27.70761430,41.91964710
2858,Jizan,194,SA,"Saudi Arabia",09,region,17.17381760,42.70761070
2850,Makkah,194,SA,"Saudi Arabia",02,region,21.52355840,41.91964710
2860,Najran,194,SA,"Saudi Arabia",10,region,18.35146640,45.60071080
2854,"Northern Borders",194,SA,"Saudi Arabia",08,region,30.07991620,42.86378750
2849,Riyadh,194,SA,"Saudi Arabia",01,region,22.75543850,46.20915470
2852,Tabuk,194,SA,"Saudi Arabia",07,region,28.24533350,37.63866220
473,Dakar,195,SN,Senegal,DK,,14.71667700,-17.46768610
480,"Diourbel Region",195,SN,Senegal,DB,,14.72830850,-16.25221430
479,Fatick,195,SN,Senegal,FK,,14.33901670,-16.41114250
475,Kaffrine,195,SN,Senegal,KA,,14.10520200,-15.54157550
483,Kaolack,195,SN,Senegal,KL,,14.16520830,-16.07577490
481,Kédougou,195,SN,Senegal,KE,,12.56046070,-12.17470770
474,Kolda,195,SN,Senegal,KD,,12.91074950,-14.95056710
485,Louga,195,SN,Senegal,LG,,15.61417680,-16.22868000
476,Matam,195,SN,Senegal,MT,,15.66002250,-13.25769060
477,Saint-Louis,195,SN,Senegal,SL,,38.62700250,-90.19940420
482,Sédhiou,195,SN,Senegal,SE,,12.70460400,-15.55623040
486,"Tambacounda Region",195,SN,Senegal,TC,,13.56190110,-13.17403480
484,"Thiès Region",195,SN,Senegal,TH,,14.79100520,-16.93586040
478,Ziguinchor,195,SN,Senegal,ZG,,12.56414790,-16.26398250
3728,Belgrade,196,RS,Serbia,00,,44.78656800,20.44892160
3717,"Bor District",196,RS,Serbia,14,,44.06989180,22.09850860
3732,"Braničevo District",196,RS,Serbia,11,,44.69822460,21.54467750
3716,"Central Banat District",196,RS,Serbia,02,,45.47884850,20.60825220
3715,"Jablanica District",196,RS,Serbia,23,,42.94815600,21.81293210
3724,"Kolubara District",196,RS,Serbia,09,,44.35098110,20.00043050
3719,"Mačva District",196,RS,Serbia,08,,44.59253140,19.50822460
3727,"Moravica District",196,RS,Serbia,17,,43.84147000,20.29049870
3722,"Nišava District",196,RS,Serbia,20,,43.37389020,21.93223310
3714,"North Bačka District",196,RS,Serbia,01,,45.98033940,19.59070010
3736,"North Banat District",196,RS,Serbia,03,,45.90683900,19.99934170
3721,"Pčinja District",196,RS,Serbia,24,,42.58363620,22.14302150
3712,"Pirot District",196,RS,Serbia,22,,43.08740360,22.59830440
3741,"Podunavlje District",196,RS,Serbia,10,,44.47291560,20.99014260
3737,"Pomoravlje District",196,RS,Serbia,13,,43.95913790,21.27135300
3720,"Rasina District",196,RS,Serbia,19,,43.52635250,21.15881780
3725,"Raška District",196,RS,Serbia,18,,43.33734610,20.57340050
3711,"South Bačka District",196,RS,Serbia,06,,45.48903440,19.69761870
3713,"South Banat District",196,RS,Serbia,04,,45.00274570,21.05425090
3740,"Srem District",196,RS,Serbia,07,,45.00291710,19.80137730
3734,"Šumadija District",196,RS,Serbia,12,,44.20506780,20.78565650
3718,"Toplica District",196,RS,Serbia,21,,43.19065920,21.34077620
3733,Vojvodina,196,RS,Serbia,VO,,45.26086510,19.83193380
3726,"West Bačka District",196,RS,Serbia,05,,45.73553850,19.18973640
3731,"Zaječar District",196,RS,Serbia,15,,43.90150480,22.27380110
3729,"Zlatibor District",196,RS,Serbia,16,,43.64541700,19.71014550
513,"Anse Boileau",197,SC,Seychelles,02,,-4.70472680,55.48593630
502,"Anse Royale",197,SC,Seychelles,05,,-4.74079880,55.50810120
506,Anse-aux-Pins,197,SC,Seychelles,01,,-4.69004430,55.51502890
508,"Au Cap",197,SC,Seychelles,04,,-4.70597230,55.50810120
497,"Baie Lazare",197,SC,Seychelles,06,,-4.74825250,55.48593630
514,"Baie Sainte Anne",197,SC,Seychelles,07,,47.05259000,-64.95245790
512,"Beau Vallon",197,SC,Seychelles,08,,-4.62109670,55.42778020
515,"Bel Air",197,SC,Seychelles,09,,34.10024550,-118.45946300
505,"Bel Ombre",197,SC,Seychelles,10,,-20.50100950,57.42596240
517,Cascade,197,SC,Seychelles,11,,44.51628210,-116.04179830
503,Glacis,197,SC,Seychelles,12,,47.11573030,-70.30281830
500,"Grand'Anse Mahé",197,SC,Seychelles,13,,-4.67739200,55.46377700
504,"Grand'Anse Praslin",197,SC,Seychelles,14,,-4.31762190,55.70783630
495,"La Digue",197,SC,Seychelles,15,,49.76669220,-97.15466290
516,"La Rivière Anglaise",197,SC,Seychelles,16,,-4.61061500,55.45408410
499,"Les Mamelles",197,SC,Seychelles,24,,38.82505050,-90.48345170
494,"Mont Buxton",197,SC,Seychelles,17,,-4.61666670,55.44577680
498,"Mont Fleuri",197,SC,Seychelles,18,,-4.63565430,55.45546880
511,Plaisance,197,SC,Seychelles,19,,45.60709500,-75.11427450
510,"Pointe La Rue",197,SC,Seychelles,20,,-4.68048900,55.51918570
507,"Port Glaud",197,SC,Seychelles,21,,-4.64885230,55.41947530
501,"Roche Caiman",197,SC,Seychelles,25,,-4.63960280,55.46793150
496,"Saint Louis",197,SC,Seychelles,22,,38.62700250,-90.19940420
509,Takamaka,197,SC,Seychelles,23,,37.96459170,-1.22177270
914,"Eastern Province",198,SL,"Sierra Leone",E,,,
911,"Northern Province",198,SL,"Sierra Leone",N,,,
912,"Southern Province",198,SL,"Sierra Leone",S,,,
913,"Western Area",198,SL,"Sierra Leone",W,,40.25459690,-80.24554440
4651,"Central Singapore",199,SG,Singapore,01,district,1.28840000,103.85350000
4649,"North East",199,SG,Singapore,02,district,1.38240000,103.89720000
4653,"North West",199,SG,Singapore,03,district,1.41800000,103.82750000
4650,"South East",199,SG,Singapore,04,district,1.35710000,103.70040000
4652,"South West",199,SG,Singapore,05,district,1.35710000,103.94510000
4352,"Banská Bystrica Region",200,SK,Slovakia,BC,,48.53124990,19.38287400
4356,"Bratislava Region",200,SK,Slovakia,BL,,48.31183040,17.19732990
4353,"Košice Region",200,SK,Slovakia,KI,,48.63757370,21.08342250
4357,"Nitra Region",200,SK,Slovakia,NI,,48.01437650,18.54165040
4354,"Prešov Region",200,SK,Slovakia,PV,,49.17167730,21.37420010
4358,"Trenčín Region",200,SK,Slovakia,TC,,48.80867580,18.23240260
4355,"Trnava Region",200,SK,Slovakia,TA,,48.39438980,17.72162050
4359,"Žilina Region",200,SK,Slovakia,ZI,,49.20314350,19.36457330
4183,"Ajdovščina Municipality",201,SI,Slovenia,001,,45.88707760,13.90428180
4326,"Ankaran Municipality",201,SI,Slovenia,213,,45.57845100,13.73691740
4301,"Beltinci Municipality",201,SI,Slovenia,002,,46.60791530,16.23651270
4166,"Benedikt Municipality",201,SI,Slovenia,148,,46.61558410,15.89572810
4179,"Bistrica ob Sotli Municipality",201,SI,Slovenia,149,,46.05655790,15.66259470
4202,"Bled Municipality",201,SI,Slovenia,003,,46.36832660,14.11457980
4278,"Bloke Municipality",201,SI,Slovenia,150,,45.77281410,14.50634590
4282,"Bohinj Municipality",201,SI,Slovenia,004,,46.30056520,13.94271950
4200,"Borovnica Municipality",201,SI,Slovenia,005,,45.90445250,14.38241890
4181,"Bovec Municipality",201,SI,Slovenia,006,,46.33804950,13.55241740
4141,"Braslovče Municipality",201,SI,Slovenia,151,,46.28361920,15.04183200
4240,"Brda Municipality",201,SI,Slovenia,007,,45.99756520,13.52704740
4215,"Brežice Municipality",201,SI,Slovenia,009,,45.90410960,15.59436390
4165,"Brezovica Municipality",201,SI,Slovenia,008,,45.95593510,14.43499520
4147,"Cankova Municipality",201,SI,Slovenia,152,,46.71823700,16.01972220
4310,"Cerklje na Gorenjskem Municipality",201,SI,Slovenia,012,,46.25170540,14.48579790
4162,"Cerknica Municipality",201,SI,Slovenia,013,,45.79662550,14.39217700
4178,"Cerkno Municipality",201,SI,Slovenia,014,,46.12884140,13.98940270
4176,"Cerkvenjak Municipality",201,SI,Slovenia,153,,46.56707110,15.94297530
4191,"City Municipality of Celje",201,SI,Slovenia,011,,46.23974950,15.26770630
4236,"City Municipality of Novo Mesto",201,SI,Slovenia,085,,45.80108240,15.17100890
4151,"Črenšovci Municipality",201,SI,Slovenia,015,,46.57200290,16.28773460
4232,"Črna na Koroškem Municipality",201,SI,Slovenia,016,,46.47045290,14.84999980
4291,"Črnomelj Municipality",201,SI,Slovenia,017,,45.53612250,15.19441430
4304,"Destrnik Municipality",201,SI,Slovenia,018,,46.49223680,15.87779560
4167,"Divača Municipality",201,SI,Slovenia,019,,45.68060690,13.97203120
4295,"Dobje Municipality",201,SI,Slovenia,154,,46.13700370,15.39412900
4216,"Dobrepolje Municipality",201,SI,Slovenia,020,,45.85249510,14.70831090
4252,"Dobrna Municipality",201,SI,Slovenia,155,,46.33561410,15.22597320
4308,"Dobrova–Polhov Gradec Municipality",201,SI,Slovenia,021,,46.06488960,14.31681950
4189,"Dobrovnik Municipality",201,SI,Slovenia,156,,46.65386620,16.35065940
4173,"Dol pri Ljubljani Municipality",201,SI,Slovenia,022,,46.08843860,14.64247920
4281,"Dolenjske Toplice Municipality",201,SI,Slovenia,157,,45.73457110,15.01294930
4159,"Domžale Municipality",201,SI,Slovenia,023,,46.14382690,14.63752790
4290,"Dornava Municipality",201,SI,Slovenia,024,,46.44435130,15.98891590
4345,"Dravograd Municipality",201,SI,Slovenia,025,,46.58921900,15.02460210
4213,"Duplek Municipality",201,SI,Slovenia,026,,46.50100160,15.75463070
4293,"Gorenja Vas–Poljane Municipality",201,SI,Slovenia,027,,46.11165820,14.11493480
4210,"Gorišnica Municipality",201,SI,Slovenia,028,,46.41202710,16.01330890
4284,"Gorje Municipality",201,SI,Slovenia,207,,46.38024580,14.06853390
4343,"Gornja Radgona Municipality",201,SI,Slovenia,029,,46.67670990,15.99108470
4339,"Gornji Grad Municipality",201,SI,Slovenia,030,,46.29617120,14.80623470
4271,"Gornji Petrovci Municipality",201,SI,Slovenia,031,,46.80371280,16.21913790
4217,"Grad Municipality",201,SI,Slovenia,158,,46.80873200,16.10920600
4336,"Grosuplje Municipality",201,SI,Slovenia,032,,45.95576450,14.65889900
4145,"Hajdina Municipality",201,SI,Slovenia,159,,46.41850140,15.82447220
4175,"Hoče–Slivnica Municipality",201,SI,Slovenia,160,,46.47785800,15.64760050
4327,"Hodoš Municipality",201,SI,Slovenia,161,,46.83141340,16.32106800
4193,"Horjul Municipality",201,SI,Slovenia,162,,46.02253780,14.29862690
4341,"Hrastnik Municipality",201,SI,Slovenia,034,,46.14172880,15.08448940
4321,"Hrpelje–Kozina Municipality",201,SI,Slovenia,035,,45.60911920,13.93791480
4152,"Idrija Municipality",201,SI,Slovenia,036,,46.00409390,13.97754930
4286,"Ig Municipality",201,SI,Slovenia,037,,45.95888680,14.52705280
4305,"Ivančna Gorica Municipality",201,SI,Slovenia,039,,45.93958410,14.80476260
4322,"Izola Municipality",201,SI,Slovenia,040,,45.53135570,13.66646490
4337,"Jesenice Municipality",201,SI,Slovenia,041,,46.43670470,14.05260570
4203,"Jezersko Municipality",201,SI,Slovenia,163,,46.39427940,14.49855590
4266,"Juršinci Municipality",201,SI,Slovenia,042,,46.48986510,15.98092300
4180,"Kamnik Municipality",201,SI,Slovenia,043,,46.22216660,14.60707270
4227,"Kanal ob Soči Municipality",201,SI,Slovenia,044,,46.06735300,13.62033500
4150,"Kidričevo Municipality",201,SI,Slovenia,045,,46.39575720,15.79259060
4243,"Kobarid Municipality",201,SI,Slovenia,046,,46.24569710,13.57869490
4325,"Kobilje Municipality",201,SI,Slovenia,047,,46.68518000,16.39367190
4335,"Kočevje Municipality",201,SI,Slovenia,048,,45.64280000,14.86158380
4315,"Komen Municipality",201,SI,Slovenia,049,,45.81752350,13.74827110
4283,"Komenda Municipality",201,SI,Slovenia,164,,46.20648800,14.53824990
4319,"Koper City Municipality",201,SI,Slovenia,050,,45.54805900,13.73018770
4254,"Kostanjevica na Krki Municipality",201,SI,Slovenia,197,,45.83166380,15.44119060
4331,"Kostel Municipality",201,SI,Slovenia,165,,45.49282550,14.87082350
4186,"Kozje Municipality",201,SI,Slovenia,051,,46.07332110,15.55967190
4287,"Kranj City Municipality",201,SI,Slovenia,052,,46.25850210,14.35435690
4340,"Kranjska Gora Municipality",201,SI,Slovenia,053,,46.48452930,13.78571450
4238,"Križevci Municipality",201,SI,Slovenia,166,,46.57018210,16.10926530
4197,Kungota,201,SI,Slovenia,055,,46.64187930,15.60362880
4211,"Kuzma Municipality",201,SI,Slovenia,056,,46.83510380,16.08071000
4338,"Laško Municipality",201,SI,Slovenia,057,,46.15422360,15.23614910
4142,"Lenart Municipality",201,SI,Slovenia,058,,46.58344240,15.82621250
4225,"Lendava Municipality",201,SI,Slovenia,059,,46.55134830,16.44198390
4347,"Litija Municipality",201,SI,Slovenia,060,,46.05732260,14.83096360
4270,"Ljubljana City Municipality",201,SI,Slovenia,061,,46.05694650,14.50575150
4294,"Ljubno Municipality",201,SI,Slovenia,062,,46.34431250,14.83354920
4351,"Ljutomer Municipality",201,SI,Slovenia,063,,46.51908480,16.18932160
4306,"Log–Dragomer Municipality",201,SI,Slovenia,208,,46.01787470,14.36877670
4350,"Logatec Municipality",201,SI,Slovenia,064,,45.91761100,14.23514510
4174,"Loška Dolina Municipality",201,SI,Slovenia,065,,45.64779080,14.49731470
4158,"Loški Potok Municipality",201,SI,Slovenia,066,,45.69096370,14.59859700
4156,"Lovrenc na Pohorju Municipality",201,SI,Slovenia,167,,46.54196380,15.40004430
4219,"Luče Municipality",201,SI,Slovenia,067,,46.35449250,14.74715040
4302,"Lukovica Municipality",201,SI,Slovenia,068,,46.16962930,14.69072590
4157,"Majšperk Municipality",201,SI,Slovenia,069,,46.35030190,15.73405950
4224,"Makole Municipality",201,SI,Slovenia,198,,46.31686970,15.66641260
4242,"Maribor City Municipality",201,SI,Slovenia,070,,46.55064960,15.62054390
4244,"Markovci Municipality",201,SI,Slovenia,168,,46.38793090,15.95860140
4349,"Medvode Municipality",201,SI,Slovenia,071,,46.14190800,14.40325960
4348,"Mengeš Municipality",201,SI,Slovenia,072,,46.16591220,14.57196940
4323,"Metlika Municipality",201,SI,Slovenia,073,,45.64807150,15.31778380
4265,"Mežica Municipality",201,SI,Slovenia,074,,46.52150270,14.85213400
4223,"Miklavž na Dravskem Polju Municipality",201,SI,Slovenia,169,,46.50826280,15.69520650
4220,"Miren–Kostanjevica Municipality",201,SI,Slovenia,075,,45.84360290,13.62766470
4298,"Mirna Municipality",201,SI,Slovenia,212,,45.95156470,15.06209770
4237,"Mirna Peč Municipality",201,SI,Slovenia,170,,45.84815740,15.08794500
4212,"Mislinja Municipality",201,SI,Slovenia,076,,46.44294030,15.19876780
4297,"Mokronog–Trebelno Municipality",201,SI,Slovenia,199,,45.90885290,15.15967360
4168,"Moravče Municipality",201,SI,Slovenia,077,,46.13627810,14.74600100
4218,"Moravske Toplice Municipality",201,SI,Slovenia,078,,46.68569320,16.22245820
4190,"Mozirje Municipality",201,SI,Slovenia,079,,46.33943500,14.96024130
4318,"Municipality of Apače",201,SI,Slovenia,195,,46.69746790,15.91025340
4309,"Municipality of Cirkulane",201,SI,Slovenia,196,,46.32983220,15.99806660
4344,"Municipality of Ilirska Bistrica",201,SI,Slovenia,038,,45.57913230,14.28097290
4314,"Municipality of Krško",201,SI,Slovenia,054,,45.95896090,15.49235550
4187,"Municipality of Škofljica",201,SI,Slovenia,123,,45.98409620,14.57466260
4313,"Murska Sobota City Municipality",201,SI,Slovenia,080,,46.64321470,16.15157540
4208,"Muta Municipality",201,SI,Slovenia,081,,46.60973660,15.16299950
4177,"Naklo Municipality",201,SI,Slovenia,082,,46.27186630,14.31569320
4329,"Nazarje Municipality",201,SI,Slovenia,083,,46.28217410,14.92256290
4205,"Nova Gorica City Municipality",201,SI,Slovenia,084,,45.97627600,13.73088810
4320,"Odranci Municipality",201,SI,Slovenia,086,,46.59010170,16.27881650
4143,Oplotnica,201,SI,Slovenia,171,,46.38716300,15.44581310
4221,"Ormož Municipality",201,SI,Slovenia,087,,46.43533330,16.15437400
4199,"Osilnica Municipality",201,SI,Slovenia,088,,45.54184670,14.71563030
4172,"Pesnica Municipality",201,SI,Slovenia,089,,46.60887550,15.67570510
4201,"Piran Municipality",201,SI,Slovenia,090,,45.52888560,13.56807350
4184,"Pivka Municipality",201,SI,Slovenia,091,,45.67892960,14.25426890
4146,"Podčetrtek Municipality",201,SI,Slovenia,092,,46.17395420,15.60138160
4161,"Podlehnik Municipality",201,SI,Slovenia,172,,46.33107820,15.87858360
4234,"Podvelka Municipality",201,SI,Slovenia,093,,46.62219520,15.38899220
4239,"Poljčane Municipality",201,SI,Slovenia,200,,46.31398530,15.57847910
4272,"Polzela Municipality",201,SI,Slovenia,173,,46.28089700,15.07373210
4330,"Postojna Municipality",201,SI,Slovenia,094,,45.77493900,14.21342630
4188,"Prebold Municipality",201,SI,Slovenia,174,,46.23591360,15.09369120
4303,"Preddvor Municipality",201,SI,Slovenia,095,,46.30171390,14.42181650
4274,"Prevalje Municipality",201,SI,Slovenia,175,,46.56211460,14.88478610
4228,"Ptuj City Municipality",201,SI,Slovenia,096,,46.41995350,15.86968840
4288,"Puconci Municipality",201,SI,Slovenia,097,,46.72004180,16.09977920
4204,"Rače–Fram Municipality",201,SI,Slovenia,098,,46.45420830,15.63294670
4195,"Radeče Municipality",201,SI,Slovenia,099,,46.06669540,15.18204380
4292,"Radenci Municipality",201,SI,Slovenia,100,,46.62311210,16.05069030
4275,"Radlje ob Dravi Municipality",201,SI,Slovenia,101,,46.61357320,15.23544380
4231,"Radovljica Municipality",201,SI,Slovenia,102,,46.33558270,14.20945340
4155,"Ravne na Koroškem Municipality",201,SI,Slovenia,103,,46.55211940,14.95990840
4206,"Razkrižje Municipality",201,SI,Slovenia,176,,46.52263390,16.26686380
4160,"Rečica ob Savinji Municipality",201,SI,Slovenia,209,,46.32337900,14.92236700
4253,"Renče–Vogrsko Municipality",201,SI,Slovenia,201,,45.89546170,13.67856730
4235,"Ribnica Municipality",201,SI,Slovenia,104,,45.74003030,14.72657820
4207,"Ribnica na Pohorju Municipality",201,SI,Slovenia,177,,46.53561450,15.26745380
4233,"Rogaška Slatina Municipality",201,SI,Slovenia,106,,46.24539730,15.62650140
4264,"Rogašovci Municipality",201,SI,Slovenia,105,,46.80557850,16.03452370
4209,"Rogatec Municipality",201,SI,Slovenia,107,,46.22866260,15.69913380
4280,"Ruše Municipality",201,SI,Slovenia,108,,46.52062650,15.48178690
4222,"Šalovci Municipality",201,SI,Slovenia,033,,46.85335680,16.25917910
4230,"Selnica ob Dravi Municipality",201,SI,Slovenia,178,,46.55139180,15.49294100
4346,"Semič Municipality",201,SI,Slovenia,109,,45.65205340,15.18207010
4317,"Šempeter–Vrtojba Municipality",201,SI,Slovenia,183,,45.92900950,13.64155940
4299,"Šenčur Municipality",201,SI,Slovenia,117,,46.24336990,14.41922230
4324,"Šentilj Municipality",201,SI,Slovenia,118,,46.68628390,15.71035670
4241,"Šentjernej Municipality",201,SI,Slovenia,119,,45.84341300,15.33783120
4171,"Šentjur Municipality",201,SI,Slovenia,120,,46.26543390,15.40800000
4311,"Šentrupert Municipality",201,SI,Slovenia,211,,45.98731420,15.08297830
4268,"Sevnica Municipality",201,SI,Slovenia,110,,46.00703170,15.30456790
4149,"Sežana Municipality",201,SI,Slovenia,111,,45.72751090,13.86619310
4170,"Škocjan Municipality",201,SI,Slovenia,121,,45.91754540,15.31017360
4316,"Škofja Loka Municipality",201,SI,Slovenia,122,,46.14098440,14.28118730
4169,"Slovenj Gradec City Municipality",201,SI,Slovenia,112,,46.48777180,15.07294780
4332,"Slovenska Bistrica Municipality",201,SI,Slovenia,113,,46.39198130,15.57278690
4198,"Slovenske Konjice Municipality",201,SI,Slovenia,114,,46.33691910,15.42147080
4285,"Šmarje pri Jelšah Municipality",201,SI,Slovenia,124,,46.22870250,15.51903530
4289,"Šmarješke Toplice Municipality",201,SI,Slovenia,206,,45.86803770,15.23474220
4296,"Šmartno ob Paki Municipality",201,SI,Slovenia,125,,46.32903720,15.03339370
4279,"Šmartno pri Litiji Municipality",201,SI,Slovenia,194,,46.04549710,14.84101330
4277,"Sodražica Municipality",201,SI,Slovenia,179,,45.76165650,14.63528530
4261,"Solčava Municipality",201,SI,Slovenia,180,,46.40235260,14.68023040
4248,"Šoštanj Municipality",201,SI,Slovenia,126,,46.37828360,15.04613780
4263,"Središče ob Dravi",201,SI,Slovenia,202,,46.39592820,16.27049150
4259,"Starše Municipality",201,SI,Slovenia,115,,46.46743310,15.76405460
4185,"Štore Municipality",201,SI,Slovenia,127,,46.22225140,15.31261160
4333,"Straža Municipality",201,SI,Slovenia,203,,45.77684280,15.09486940
4164,"Sveta Ana Municipality",201,SI,Slovenia,181,,46.65000000,15.84527800
4260,"Sveta Trojica v Slovenskih Goricah Municipality",201,SI,Slovenia,204,,46.56808090,15.88230640
4229,"Sveti Andraž v Slovenskih Goricah Municipality",201,SI,Slovenia,182,,46.51897470,15.94982620
4255,"Sveti Jurij ob Ščavnici Municipality",201,SI,Slovenia,116,,46.56874520,16.02225280
4328,"Sveti Jurij v Slovenskih Goricah Municipality",201,SI,Slovenia,210,,46.61707910,15.78046770
4273,"Sveti Tomaž Municipality",201,SI,Slovenia,205,,46.48352830,16.07944200
4194,"Tabor Municipality",201,SI,Slovenia,184,,46.21079210,15.01742490
4312,"Tišina Municipality",201,SI,Slovenia,010,,46.65418840,16.07547810
4247,"Tolmin Municipality",201,SI,Slovenia,128,,46.18571880,13.73198380
4246,"Trbovlje Municipality",201,SI,Slovenia,129,,46.15035630,15.04531370
4214,"Trebnje Municipality",201,SI,Slovenia,130,,45.90801630,15.01319050
4153,"Trnovska Vas Municipality",201,SI,Slovenia,185,,46.52940350,15.88531180
4250,"Tržič Municipality",201,SI,Slovenia,131,,46.35935140,14.30066230
4334,"Trzin Municipality",201,SI,Slovenia,186,,46.12982410,14.55776370
4251,"Turnišče Municipality",201,SI,Slovenia,132,,46.61375040,16.32021000
4267,"Velika Polana Municipality",201,SI,Slovenia,187,,46.57317150,16.34441260
4144,"Velike Lašče Municipality",201,SI,Slovenia,134,,45.83365910,14.63623630
4257,"Veržej Municipality",201,SI,Slovenia,188,,46.58411350,16.16208000
4300,"Videm Municipality",201,SI,Slovenia,135,,46.36383300,15.87812120
4196,"Vipava Municipality",201,SI,Slovenia,136,,45.84126740,13.96096130
4148,"Vitanje Municipality",201,SI,Slovenia,137,,46.38153230,15.29506870
4154,"Vodice Municipality",201,SI,Slovenia,138,,46.18966430,14.49385390
4245,"Vojnik Municipality",201,SI,Slovenia,139,,46.29205810,15.30205800
4163,"Vransko Municipality",201,SI,Slovenia,189,,46.23900600,14.95272490
4262,"Vrhnika Municipality",201,SI,Slovenia,140,,45.95027190,14.32764220
4226,"Vuzenica Municipality",201,SI,Slovenia,141,,46.59808360,15.16572370
4269,"Zagorje ob Savi Municipality",201,SI,Slovenia,142,,46.13452020,14.99643840
4258,"Žalec Municipality",201,SI,Slovenia,190,,46.25197120,15.16500720
4182,"Zavrč Municipality",201,SI,Slovenia,143,,46.35713000,16.04777470
4256,"Železniki Municipality",201,SI,Slovenia,146,,46.22563770,14.16936170
4249,"Žetale Municipality",201,SI,Slovenia,191,,46.27428330,15.79133590
4192,"Žiri Municipality",201,SI,Slovenia,147,,46.04724990,14.10984510
4276,"Žirovnica Municipality",201,SI,Slovenia,192,,46.39544030,14.15396320
4342,"Zreče Municipality",201,SI,Slovenia,144,,46.41777860,15.37094310
4307,"Žužemberk Municipality",201,SI,Slovenia,193,,45.82003500,14.95359190
4784,"Central Province",202,SB,"Solomon Islands",CE,,,
4781,"Choiseul Province",202,SB,"Solomon Islands",CH,,-7.05014940,156.95114590
4785,"Guadalcanal Province",202,SB,"Solomon Islands",GU,,-9.57732840,160.14558050
4778,Honiara,202,SB,"Solomon Islands",CT,,-9.44563810,159.97289990
4780,"Isabel Province",202,SB,"Solomon Islands",IS,,-8.05923530,159.14470810
4782,"Makira-Ulawa Province",202,SB,"Solomon Islands",MK,,-10.57374470,161.80969410
4783,"Malaita Province",202,SB,"Solomon Islands",ML,,-8.94461680,160.90712360
4787,"Rennell and Bellona Province",202,SB,"Solomon Islands",RB,,-11.61314350,160.16939490
4779,"Temotu Province",202,SB,"Solomon Islands",TE,,-10.68692900,166.06239790
4786,"Western Province",202,SB,"Solomon Islands",WE,,,
925,"Awdal Region",203,SO,Somalia,AW,,10.63342850,43.32946600
917,Bakool,203,SO,Somalia,BK,,4.36572210,44.09603110
927,Banaadir,203,SO,Somalia,BN,,2.11873750,45.33694590
930,Bari,203,SO,Somalia,BR,,41.11714320,16.87187150
926,Bay,203,SO,Somalia,BY,,37.03655340,-95.61747670
918,Galguduud,203,SO,Somalia,GA,,5.18501280,46.82528380
928,Gedo,203,SO,Somalia,GE,,3.50392270,42.23624350
915,Hiran,203,SO,Somalia,HI,,4.32101500,45.29938620
924,"Lower Juba",203,SO,Somalia,JH,,0.22402100,41.60118140
921,"Lower Shebelle",203,SO,Somalia,SH,,1.87664580,44.24790150
922,"Middle Juba",203,SO,Somalia,JD,,2.07804880,41.60118140
923,"Middle Shebelle",203,SO,Somalia,SD,,2.92502470,45.90396890
916,Mudug,203,SO,Somalia,MU,,6.56567260,47.76375650
920,Nugal,203,SO,Somalia,NU,,43.27938610,17.03392050
919,"Sanaag Region",203,SO,Somalia,SA,,10.39382180,47.76375650
929,"Togdheer Region",203,SO,Somalia,TO,,9.44605870,45.29938620
938,"Eastern Cape",204,ZA,"South Africa",EC,,-32.29684020,26.41938900
932,"Free State",204,ZA,"South Africa",FS,,37.68585250,-97.28112560
936,Gauteng,204,ZA,"South Africa",GP,,-26.27075930,28.11226790
935,KwaZulu-Natal,204,ZA,"South Africa",KZN,,-28.53055390,30.89582420
933,Limpopo,204,ZA,"South Africa",LP,,-23.40129460,29.41793240
937,Mpumalanga,204,ZA,"South Africa",MP,,-25.56533600,30.52790960
934,"North West",204,ZA,"South Africa",NW,,32.75885200,-97.32880600
931,"Northern Cape",204,ZA,"South Africa",NC,,-29.04668080,21.85685860
939,"Western Cape",204,ZA,"South Africa",WC,,-33.22779180,21.85685860
3860,Busan,116,KR,"South Korea",26,,35.17955430,129.07564160
3846,Daegu,116,KR,"South Korea",27,,35.87143540,128.60144500
3850,Daejeon,116,KR,"South Korea",30,,36.35041190,127.38454750
3862,"Gangwon Province",116,KR,"South Korea",42,,37.82280000,128.15550000
3858,Gwangju,116,KR,"South Korea",29,,35.15954540,126.85260120
3847,"Gyeonggi Province",116,KR,"South Korea",41,,37.41380000,127.51830000
3848,Incheon,116,KR,"South Korea",28,,37.45625570,126.70520620
3853,Jeju,116,KR,"South Korea",49,,33.95682780,-84.13135000
3854,"North Chungcheong Province",116,KR,"South Korea",43,,36.80000000,127.70000000
3855,"North Gyeongsang Province",116,KR,"South Korea",47,,36.49190000,128.88890000
3851,"North Jeolla Province",116,KR,"South Korea",45,,35.71750000,127.15300000
3861,"Sejong City",116,KR,"South Korea",50,,34.05233230,-118.30848970
3849,Seoul,116,KR,"South Korea",11,,37.56653500,126.97796920
3859,"South Chungcheong Province",116,KR,"South Korea",44,,36.51840000,126.80000000
3857,"South Gyeongsang Province",116,KR,"South Korea",48,,35.46060000,128.21320000
3856,"South Jeolla Province",116,KR,"South Korea",46,,34.86790000,126.99100000
3852,Ulsan,116,KR,"South Korea",31,,35.53837730,129.31135960
2092,"Central Equatoria",206,SS,"South Sudan",EC,,4.61440630,31.26263660
2093,"Eastern Equatoria",206,SS,"South Sudan",EE,,5.06929950,33.43835300
2094,"Jonglei State",206,SS,"South Sudan",JG,,7.18196190,32.35609520
2090,Lakes,206,SS,"South Sudan",LK,,37.16282550,-95.69116230
2088,"Northern Bahr el Ghazal",206,SS,"South Sudan",BN,,8.53604490,26.79678490
2085,Unity,206,SS,"South Sudan",UY,,37.78712760,-122.40340790
2086,"Upper Nile",206,SS,"South Sudan",NU,,9.88942020,32.71813750
2087,Warrap,206,SS,"South Sudan",WR,,8.08862380,28.64106410
2091,"Western Bahr el Ghazal",206,SS,"South Sudan",BW,,8.64523990,25.28375850
2089,"Western Equatoria",206,SS,"South Sudan",EW,,5.34717990,28.29943500
5089,"A Coruña",207,ES,Spain,C,province,43.36190400,-8.43019320
5109,Albacete,207,ES,Spain,AB,province,38.99223120,-1.87809890
5108,Alicante,207,ES,Spain,A,province,38.35795460,-0.54256340
5095,Almeria,207,ES,Spain,AL,province,36.84152680,-2.47462610
5093,Araba,207,ES,Spain,VI,province,42.83951190,-3.84237740
1160,Asturias,207,ES,Spain,O,province,43.36139530,-5.85932670
1189,Ávila,207,ES,Spain,AV,province,40.69345110,-4.89356270
5092,Badajoz,207,ES,Spain,BA,province,38.87937480,-7.02269830
5102,Barcelona,207,ES,Spain,B,province,41.39266790,2.14018910
5094,Bizkaia,207,ES,Spain,BI,province,43.21921990,-3.21110870
1146,Burgos,207,ES,Spain,BU,province,42.33807580,-3.58126920
1190,Caceres,207,ES,Spain,CC,province,39.47163130,-6.42573840
5096,Cádiz,207,ES,Spain,CA,province,36.51638510,-6.29997670
5222,Canarias,207,ES,Spain,CN,"autonomous community",28.29160000,16.62910000
1170,Cantabria,207,ES,Spain,S,province,43.18283960,-3.98784270
5110,Castellón,207,ES,Spain,CS,province,39.98114350,0.00884070
5223,Ceuta,207,ES,Spain,CE,"autonomous city in North Africa",35.88900000,-5.31870000
5105,"Ciudad Real",207,ES,Spain,CR,province,38.98607580,-3.94449750
5097,Córdoba,207,ES,Spain,CO,province,36.51638510,-6.29997670
5106,Cuenca,207,ES,Spain,CU,province,40.06200360,-2.16553440
1191,Gipuzkoa,207,ES,Spain,SS,province,43.14523600,-2.44618250
5103,Girona,207,ES,Spain,GI,province,41.98034450,2.80115770
5098,Granada,207,ES,Spain,GR,province,37.18094110,-3.62629100
5107,Guadalajara,207,ES,Spain,GU,province,40.63222140,-3.19068200
5099,Huelva,207,ES,Spain,H,province,37.27086660,-6.95719990
1177,Huesca,207,ES,Spain,HU,province,41.59762750,-0.90566230
1174,"Islas Baleares",207,ES,Spain,PM,province,39.35877590,2.73563280
5100,Jaén,207,ES,Spain,J,province,37.78009310,-3.81437450
1171,"La Rioja",207,ES,Spain,LO,province,42.28707330,-2.53960300
1185,"Las Palmas",207,ES,Spain,GC,province,28.29156370,-16.62913040
1200,León,207,ES,Spain,LE,province,42.59870410,-5.56708390
5104,Lleida,207,ES,Spain,L,province,41.61837310,0.60242530
5090,Lugo,207,ES,Spain,LU,province,43.01231370,-7.57400960
1158,Madrid,207,ES,Spain,M,province,40.41675150,-3.70383220
5101,Málaga,207,ES,Spain,MA,province,36.71820150,-4.51930600
5224,Melilla,207,ES,Spain,ML,"autonomous city in North Africa",35.29370000,-2.93830000
1176,Murcia,207,ES,Spain,MU,province,38.13981410,-1.36621600
1204,Navarra,207,ES,Spain,NA,province,42.69539090,-1.67606910
5091,Ourense,207,ES,Spain,OR,province,42.33836130,-7.88119510
1157,Palencia,207,ES,Spain,P,province,42.00968320,-4.52879490
1167,Pontevedra,207,ES,Spain,PO,province,42.43385950,-8.65685520
1147,Salamanca,207,ES,Spain,SA,province,40.95152630,-6.23759470
5112,"Santa Cruz de Tenerife",207,ES,Spain,TF,province,28.45789140,-16.32135390
1192,Segovia,207,ES,Spain,SG,province,40.94292960,-4.10889420
1193,Sevilla,207,ES,Spain,SE,province,37.37535010,-6.02509730
1208,Soria,207,ES,Spain,SO,province,41.76654640,-2.47903060
1203,Tarragona,207,ES,Spain,T,province,41.12586420,1.20356420
5111,Teruel,207,ES,Spain,TE,province,40.34504100,-1.11847440
1205,Toledo,207,ES,Spain,TO,province,39.86232000,-4.06946920
1175,Valencia,207,ES,Spain,V,province,39.48401080,-0.75328090
1183,Valladolid,207,ES,Spain,VA,province,41.65173750,-4.72449500
1161,Zamora,207,ES,Spain,ZA,province,41.60957440,-5.89871390
5113,Zaragoza,207,ES,Spain,Z,province,41.65175010,-0.93000020
2799,"Ampara District",208,LK,"Sri Lanka",52,,7.29116850,81.67237610
2816,"Anuradhapura District",208,LK,"Sri Lanka",71,,8.33183050,80.40290170
2790,"Badulla District",208,LK,"Sri Lanka",81,,6.99340090,81.05498150
2818,"Batticaloa District",208,LK,"Sri Lanka",51,,7.82927810,81.47183870
2798,"Central Province",208,LK,"Sri Lanka",2,,,
2815,"Colombo District",208,LK,"Sri Lanka",11,,6.92695570,79.86173060
2808,"Eastern Province",208,LK,"Sri Lanka",5,,,
2792,"Galle District",208,LK,"Sri Lanka",31,,6.05774900,80.21755720
2804,"Gampaha District",208,LK,"Sri Lanka",12,,7.07126190,80.00877460
2791,"Hambantota District",208,LK,"Sri Lanka",33,,6.15358160,81.12714900
2787,"Jaffna District",208,LK,"Sri Lanka",41,,9.69304680,80.16518540
2789,"Kalutara District",208,LK,"Sri Lanka",13,,6.60846860,80.14285840
2788,"Kandy District",208,LK,"Sri Lanka",21,,7.29315880,80.63501070
2797,"Kegalle District",208,LK,"Sri Lanka",92,,7.12040530,80.32131060
2793,"Kilinochchi District",208,LK,"Sri Lanka",42,,9.36779710,80.32131060
2805,"Mannar District",208,LK,"Sri Lanka",43,,8.98095310,79.90439750
2810,"Matale District",208,LK,"Sri Lanka",22,,7.46596460,80.62342590
2806,"Matara District",208,LK,"Sri Lanka",32,,5.94493480,80.54879970
2819,"Monaragala District",208,LK,"Sri Lanka",82,,6.87277810,81.35068320
2814,"Mullaitivu District",208,LK,"Sri Lanka",45,,9.26753880,80.81282540
2800,"North Central Province",208,LK,"Sri Lanka",7,,8.19956380,80.63269160
2817,"North Western Province",208,LK,"Sri Lanka",6,,7.75840910,80.18750650
2813,"Northern Province",208,LK,"Sri Lanka",4,,,
2794,"Nuwara Eliya District",208,LK,"Sri Lanka",23,,6.96065320,80.76927580
2812,"Polonnaruwa District",208,LK,"Sri Lanka",72,,7.93955670,81.00034030
2796,"Puttalam District",208,LK,"Sri Lanka",62,,8.02599150,79.84712720
2807,"Ratnapura district",208,LK,"Sri Lanka",91,,6.70551680,80.38483890
2803,"Sabaragamuwa Province",208,LK,"Sri Lanka",9,,6.73959410,80.36586500
2801,"Southern Province",208,LK,"Sri Lanka",3,,,
2795,"Trincomalee District",208,LK,"Sri Lanka",53,,8.60130690,81.11960750
2811,"Uva Province",208,LK,"Sri Lanka",8,,6.84276120,81.33994140
2809,"Vavuniya District",208,LK,"Sri Lanka",44,,8.75947390,80.50003340
2802,"Western Province",208,LK,"Sri Lanka",1,,,
885,"Al Jazirah",209,SD,Sudan,GZ,,14.88596110,33.43835300
886,"Al Qadarif",209,SD,Sudan,GD,,14.02430700,35.36856790
887,"Blue Nile",209,SD,Sudan,NB,,47.59867300,-122.33441900
896,"Central Darfur",209,SD,Sudan,DC,,14.37827470,24.90422080
892,"East Darfur",209,SD,Sudan,DE,,14.37827470,24.90422080
884,Kassala,209,SD,Sudan,KA,,15.45813320,36.40396290
881,Khartoum,209,SD,Sudan,KH,,15.50065440,32.55989940
890,"North Darfur",209,SD,Sudan,DN,,15.76619690,24.90422080
893,"North Kordofan",209,SD,Sudan,KN,,13.83064410,29.41793240
895,Northern,209,SD,Sudan,NO,,38.06381700,-84.46286480
880,"Red Sea",209,SD,Sudan,RS,,20.28023200,38.51257300
891,"River Nile",209,SD,Sudan,NR,,23.97275950,32.87492060
882,Sennar,209,SD,Sudan,SI,,13.56746900,33.56720450
894,"South Darfur",209,SD,Sudan,DS,,11.64886390,24.90422080
883,"South Kordofan",209,SD,Sudan,KS,,11.19901920,29.41793240
888,"West Darfur",209,SD,Sudan,DW,,12.84635610,23.00119890
889,"West Kordofan",209,SD,Sudan,GK,,11.19901920,29.41793240
879,"White Nile",209,SD,Sudan,NW,,9.33215160,31.46153000
2846,"Brokopondo District",210,SR,Suriname,BR,,4.77102470,-55.04933750
2839,"Commewijne District",210,SR,Suriname,CM,,5.74021100,-54.87312190
2842,"Coronie District",210,SR,Suriname,CR,,5.69432710,-56.29293810
2845,"Marowijne District",210,SR,Suriname,MA,,5.62681280,-54.25931180
2840,"Nickerie District",210,SR,Suriname,NI,,5.58554690,-56.83111170
2841,"Para District",210,SR,Suriname,PR,,5.48173180,-55.22592070
2843,"Paramaribo District",210,SR,Suriname,PM,,5.85203550,-55.20382780
2848,"Saramacca District",210,SR,Suriname,SA,,5.72408130,-55.66896360
2847,"Sipaliwini District",210,SR,Suriname,SI,,3.65673820,-56.20353870
2844,"Wanica District",210,SR,Suriname,WA,,5.73237620,-55.27012350
1537,"Blekinge County",213,SE,Sweden,K,,56.28333333,15.11666666
1534,"Dalarna County",213,SE,Sweden,W,,61.09170120,14.66636530
1533,"Gävleborg County",213,SE,Sweden,X,,61.30119930,16.15342140
1546,"Gotland County",213,SE,Sweden,I,,57.46841210,18.48674470
1548,"Halland County",213,SE,Sweden,N,,56.89668050,12.80339930
5117,"Jämtland County",213,SE,Sweden,0,SE-Z,63.28306200,14.23828100
1550,"Jönköping County",213,SE,Sweden,F,,57.37084340,14.34391740
1544,"Kalmar County",213,SE,Sweden,H,,57.23501560,16.18493490
1542,"Kronoberg County",213,SE,Sweden,G,,56.71834030,14.41146730
1538,"Norrbotten County",213,SE,Sweden,BD,,66.83092160,20.39919660
1539,"Örebro County",213,SE,Sweden,T,,59.53503600,15.00657310
1536,"Östergötland County",213,SE,Sweden,E,,58.34536350,15.51978440
1541,"Skåne County",213,SE,Sweden,M,,55.99025720,13.59576920
1540,"Södermanland County",213,SE,Sweden,D,,59.03363490,16.75188990
1551,"Stockholm County",213,SE,Sweden,AB,,59.60249580,18.13843830
1545,"Uppsala County",213,SE,Sweden,C,,60.00922620,17.27145880
1535,"Värmland County",213,SE,Sweden,S,,59.72940650,13.23540240
1543,"Västerbotten County",213,SE,Sweden,AC,,65.33373110,16.51616940
1552,"Västernorrland County",213,SE,Sweden,Y,,63.42764730,17.72924440
1549,"Västmanland County",213,SE,Sweden,U,,59.67138790,16.21589530
1547,"Västra Götaland County",213,SE,Sweden,O,,58.25279260,13.05964250
1639,Aargau,214,CH,Switzerland,AG,canton,47.38766640,8.25542950
1655,"Appenzell Ausserrhoden",214,CH,Switzerland,AR,canton,47.36648100,9.30009160
1649,"Appenzell Innerrhoden",214,CH,Switzerland,AI,canton,47.31619250,9.43165730
1641,Basel-Land,214,CH,Switzerland,BL,canton,47.44181220,7.76440020
4957,Basel-Stadt,214,CH,Switzerland,BS,canton,47.56666700,7.60000000
1645,Bern,214,CH,Switzerland,BE,canton,46.79886210,7.70807010
1640,Fribourg,214,CH,Switzerland,FR,canton,46.68167480,7.11726350
1647,Geneva,214,CH,Switzerland,GE,canton,46.21800730,6.12169250
1661,Glarus,214,CH,Switzerland,GL,canton,47.04112320,9.06790000
1660,Graubünden,214,CH,Switzerland,GR,canton,46.65698710,9.57802570
1658,Jura,214,CH,Switzerland,JU,canton,47.34444740,7.14306080
1663,Lucerne,214,CH,Switzerland,LU,canton,47.07956710,8.16624450
1659,Neuchâtel,214,CH,Switzerland,NE,canton,46.98998740,6.92927320
1652,Nidwalden,214,CH,Switzerland,NW,canton,46.92670160,8.38499820
1650,Obwalden,214,CH,Switzerland,OW,canton,46.87785800,8.25124900
1654,Schaffhausen,214,CH,Switzerland,SH,canton,47.70093640,8.56800400
1653,Schwyz,214,CH,Switzerland,SZ,canton,47.02071380,8.65298840
1662,Solothurn,214,CH,Switzerland,SO,canton,47.33207170,7.63883850
1644,"St. Gallen",214,CH,Switzerland,SG,canton,47.14562540,9.35043320
1657,Thurgau,214,CH,Switzerland,TG,canton,47.60378560,9.05573710
1643,Ticino,214,CH,Switzerland,TI,canton,46.33173400,8.80045290
1642,Uri,214,CH,Switzerland,UR,canton,41.48606470,-71.53085370
1648,Valais,214,CH,Switzerland,VS,canton,46.19046140,7.54492260
1651,Vaud,214,CH,Switzerland,VD,canton,46.56131350,6.53676500
1646,Zug,214,CH,Switzerland,ZG,canton,47.16615050,8.51547490
1656,Zürich,214,CH,Switzerland,ZH,canton,47.35953600,8.63564520
2941,Al-Hasakah,215,SY,Syria,HA,,36.40551500,40.79691490
2944,Al-Raqqah,215,SY,Syria,RA,,35.95941060,38.99810520
2946,Aleppo,215,SY,Syria,HL,,36.22623930,37.46813960
2936,As-Suwayda,215,SY,Syria,SU,,32.79891560,36.78195050
2939,Damascus,215,SY,Syria,DI,,33.51514440,36.39313540
2945,Daraa,215,SY,Syria,DR,,32.92488130,36.17626150
2937,"Deir ez-Zor",215,SY,Syria,DY,,35.28797980,40.30886260
2934,Hama,215,SY,Syria,HM,,35.18878650,37.21158290
2942,Homs,215,SY,Syria,HI,,34.25671230,38.31657250
2940,Idlib,215,SY,Syria,ID,,35.82687980,36.69572160
2938,Latakia,215,SY,Syria,LA,,35.61297910,36.00232250
2943,Quneitra,215,SY,Syria,QU,,33.07763180,35.89341360
2935,"Rif Dimashq",215,SY,Syria,RD,,33.51672890,36.95410700
2947,Tartus,215,SY,Syria,TA,,35.00066520,36.00232250
3404,Changhua,216,TW,Taiwan,CHA,county,24.05179630,120.51613520
3408,Chiayi,216,TW,Taiwan,CYI,city,23.45184280,120.25546150
3418,Chiayi,216,TW,Taiwan,CYQ,county,23.48007510,120.44911130
3423,Hsinchu,216,TW,Taiwan,HSQ,county,24.83872260,121.01772460
3417,Hsinchu,216,TW,Taiwan,HSZ,city,24.81382870,120.96747980
3411,Hualien,216,TW,Taiwan,HUA,county,23.98715890,121.60157140
3412,Kaohsiung,216,TW,Taiwan,KHH,"special municipality",22.62727840,120.30143530
4965,Keelung,216,TW,Taiwan,KEE,city,25.12418620,121.64758340
3415,Kinmen,216,TW,Taiwan,KIN,county,24.34877920,118.32856440
3420,Lienchiang,216,TW,Taiwan,LIE,county,26.15055560,119.92888890
3413,Miaoli,216,TW,Taiwan,MIA,county,24.56015900,120.82142650
3407,Nantou,216,TW,Taiwan,NAN,county,23.96099810,120.97186380
4966,"New Taipei",216,TW,Taiwan,NWT,"special municipality",24.98752780,121.36459470
3403,Penghu,216,TW,Taiwan,PEN,county,23.57118990,119.57931570
3405,Pingtung,216,TW,Taiwan,PIF,county,22.55197590,120.54875970
3406,Taichung,216,TW,Taiwan,TXG,"special municipality",24.14773580,120.67364820
3421,Tainan,216,TW,Taiwan,TNN,"special municipality",22.99972810,120.22702770
3422,Taipei,216,TW,Taiwan,TPE,"special municipality",25.03296940,121.56541770
3410,Taitung,216,TW,Taiwan,TTT,county,22.79724470,121.07137020
3419,Taoyuan,216,TW,Taiwan,TAO,"special municipality",24.99362810,121.30097980
3402,Yilan,216,TW,Taiwan,ILA,county,24.70210730,121.73775020
3416,Yunlin,216,TW,Taiwan,YUN,county,23.70920330,120.43133730
3397,"districts of Republican Subordination",217,TJ,Tajikistan,RA,,39.08579020,70.24083250
3399,"Gorno-Badakhshan Autonomous Province",217,TJ,Tajikistan,GB,,38.41273200,73.08774900
3398,"Khatlon Province",217,TJ,Tajikistan,KT,,37.91135620,69.09702300
3400,"Sughd Province",217,TJ,Tajikistan,SU,,39.51553260,69.09702300
1491,Arusha,218,TZ,Tanzania,01,Region,-3.38692540,36.68299270
1490,"Dar es Salaam",218,TZ,Tanzania,02,Region,-6.79235400,39.20832840
1466,Dodoma,218,TZ,Tanzania,03,Region,-6.57382280,36.26308460
1481,Geita,218,TZ,Tanzania,27,Region,-2.82422570,32.26538870
1489,Iringa,218,TZ,Tanzania,04,Region,-7.78874420,35.56578620
1465,Kagera,218,TZ,Tanzania,05,Region,-1.30011150,31.26263660
1482,Katavi,218,TZ,Tanzania,28,Region,-6.36771250,31.26263660
1478,Kigoma,218,TZ,Tanzania,08,Region,-4.88240920,29.66150550
1467,Kilimanjaro,218,TZ,Tanzania,09,Region,-4.13369270,37.80876930
1483,Lindi,218,TZ,Tanzania,12,Region,-9.23433940,38.31657250
1484,Manyara,218,TZ,Tanzania,26,Region,-4.31500580,36.95410700
1468,Mara,218,TZ,Tanzania,13,Region,-1.77535380,34.15319470
4955,Mbeya,218,TZ,Tanzania,14,Region,-8.28661120,32.81325370
1470,Morogoro,218,TZ,Tanzania,16,Region,-8.81371730,36.95410700
1476,Mtwara,218,TZ,Tanzania,17,Region,-10.33984550,40.16574660
1479,Mwanza,218,TZ,Tanzania,18,Region,-2.46711970,32.89868120
1480,Njombe,218,TZ,Tanzania,29,Region,-9.24226320,35.12687810
1488,"Pemba North",218,TZ,Tanzania,06,Region,-5.03193520,39.77555710
1472,"Pemba South",218,TZ,Tanzania,10,Region,-5.31469610,39.75495110
1485,Pwani,218,TZ,Tanzania,19,Region,-7.32377140,38.82054540
1477,Rukwa,218,TZ,Tanzania,20,Region,-8.01094440,31.44561790
1486,Ruvuma,218,TZ,Tanzania,21,Region,-10.68787170,36.26308460
1463,Shinyanga,218,TZ,Tanzania,22,Region,-3.68099610,33.42714030
1464,Simiyu,218,TZ,Tanzania,30,Region,-2.83087380,34.15319470
1474,Singida,218,TZ,Tanzania,23,Region,-6.74533520,34.15319470
4956,Songwe,218,TZ,Tanzania,31,Region,-8.27261200,31.71131740
1469,Tabora,218,TZ,Tanzania,24,Region,-5.03421380,32.80844960
1487,Tanga,218,TZ,Tanzania,25,Region,-5.30497890,38.31657250
1473,"Zanzibar North",218,TZ,Tanzania,07,Region,-5.93950930,39.27910110
1471,"Zanzibar South",218,TZ,Tanzania,11,Region,-6.26428510,39.44502810
1475,"Zanzibar West",218,TZ,Tanzania,15,Region,-6.22981360,39.25832930
3523,"Amnat Charoen",219,TH,Thailand,37,,15.86567830,104.62577740
3519,"Ang Thong",219,TH,Thailand,15,,14.58960540,100.45505200
3554,Bangkok,219,TH,Thailand,10,,13.75633090,100.50176510
3533,"Bueng Kan",219,TH,Thailand,38,,18.36091040,103.64644630
3534,"Buri Ram",219,TH,Thailand,31,,14.99510030,103.11159150
3552,Chachoengsao,219,TH,Thailand,24,,13.69041940,101.07795960
3522,"Chai Nat",219,TH,Thailand,18,,15.18519710,100.12512500
4954,Chaiyaphum,219,TH,Thailand,36,,16.00749740,101.61291720
3486,Chanthaburi,219,TH,Thailand,22,,12.61124850,102.10378060
3491,"Chiang Mai",219,TH,Thailand,50,,18.78834390,98.98530080
3498,"Chiang Rai",219,TH,Thailand,57,,19.91047980,99.84057600
3513,"Chon Buri",219,TH,Thailand,20,,13.36114310,100.98467170
3526,Chumphon,219,TH,Thailand,86,,10.49304960,99.18001990
3550,Kalasin,219,TH,Thailand,46,,16.43850800,103.50609940
3516,"Kamphaeng Phet",219,TH,Thailand,62,,16.48277980,99.52266180
3511,Kanchanaburi,219,TH,Thailand,71,,14.10113930,99.41794310
3485,"Khon Kaen",219,TH,Thailand,40,,16.43219380,102.82362140
3478,Krabi,219,TH,Thailand,81,,8.08629970,98.90628350
3544,Lampang,219,TH,Thailand,52,,18.28553950,99.51278950
3483,Lamphun,219,TH,Thailand,51,,18.57446060,99.00872210
3509,Loei,219,TH,Thailand,42,,17.48602320,101.72230020
3543,"Lop Buri",219,TH,Thailand,16,,14.79950810,100.65337060
3505,"Mae Hong Son",219,TH,Thailand,58,,19.30202960,97.96543680
3517,"Maha Sarakham",219,TH,Thailand,44,,16.01320150,103.16151690
3546,Mukdahan,219,TH,Thailand,49,,16.54359140,104.70241210
3535,"Nakhon Nayok",219,TH,Thailand,26,,14.20694660,101.21305110
3503,"Nakhon Pathom",219,TH,Thailand,73,,13.81402930,100.03729290
3548,"Nakhon Phanom",219,TH,Thailand,48,,17.39203900,104.76955080
3497,"Nakhon Ratchasima",219,TH,Thailand,30,,14.97384930,102.08365200
3492,"Nakhon Sawan",219,TH,Thailand,60,,15.69873820,100.11996000
3520,"Nakhon Si Thammarat",219,TH,Thailand,80,,8.43248310,99.95990330
3530,Nan,219,TH,Thailand,55,,45.52220800,-122.98632810
3553,Narathiwat,219,TH,Thailand,96,,6.42546070,101.82531430
3480,"Nong Bua Lam Phu",219,TH,Thailand,39,,17.22182470,102.42603680
3484,"Nong Khai",219,TH,Thailand,43,,17.87828030,102.74126380
3495,Nonthaburi,219,TH,Thailand,12,,13.85910840,100.52165080
3500,"Pathum Thani",219,TH,Thailand,13,,14.02083910,100.52502760
3540,Pattani,219,TH,Thailand,94,,6.76183080,101.32325490
3507,Pattaya,219,TH,Thailand,S,,12.92355570,100.88245510
3549,Phangnga,219,TH,Thailand,82,,8.45014140,98.52553170
3488,Phatthalung,219,TH,Thailand,93,,7.61668230,100.07402310
3538,Phayao,219,TH,Thailand,56,,19.21543670,100.20236920
3515,Phetchabun,219,TH,Thailand,67,,16.30166900,101.11928040
3532,Phetchaburi,219,TH,Thailand,76,,12.96492150,99.64258830
3514,Phichit,219,TH,Thailand,66,,16.27408760,100.33469910
3506,Phitsanulok,219,TH,Thailand,65,,16.82112380,100.26585160
3494,"Phra Nakhon Si Ayutthaya",219,TH,Thailand,14,,14.36923250,100.58766340
3528,Phrae,219,TH,Thailand,54,,18.14457740,100.14028310
3536,Phuket,219,TH,Thailand,83,,7.88044790,98.39225040
3542,"Prachin Buri",219,TH,Thailand,25,,14.04206990,101.66008740
3508,"Prachuap Khiri Khan",219,TH,Thailand,77,,11.79383890,99.79575640
3479,Ranong,219,TH,Thailand,85,,9.95287020,98.60846410
3499,Ratchaburi,219,TH,Thailand,70,,13.52828930,99.81342110
3518,Rayong,219,TH,Thailand,21,,12.68139570,101.28162610
3510,"Roi Et",219,TH,Thailand,45,,16.05381960,103.65200360
3529,"Sa Kaeo",219,TH,Thailand,27,,13.82403800,102.06458390
3501,"Sakon Nakhon",219,TH,Thailand,47,,17.16642110,104.14860550
3481,"Samut Prakan",219,TH,Thailand,11,,13.59909610,100.59983190
3504,"Samut Sakhon",219,TH,Thailand,74,,13.54752160,100.27439560
3502,"Samut Songkhram",219,TH,Thailand,75,,13.40982170,100.00226450
3487,Saraburi,219,TH,Thailand,19,,14.52891540,100.91014210
3537,Satun,219,TH,Thailand,91,,6.62381580,100.06737440
3547,"Si Sa Ket",219,TH,Thailand,33,,15.11860090,104.32200950
3490,"Sing Buri",219,TH,Thailand,17,,14.89362530,100.39673140
3539,Songkhla,219,TH,Thailand,90,,7.18976590,100.59538130
3545,Sukhothai,219,TH,Thailand,64,,43.64855560,-79.37466390
3524,"Suphan Buri",219,TH,Thailand,72,,14.47448920,100.11771280
3482,"Surat Thani",219,TH,Thailand,84,,9.13419490,99.33341980
3531,Surin,219,TH,Thailand,32,,37.03582710,-95.62763670
3525,Tak,219,TH,Thailand,63,,45.02996460,-93.10498150
3541,Trang,219,TH,Thailand,92,,7.56448330,99.62393340
3496,Trat,219,TH,Thailand,23,,12.24275630,102.51747340
3512,"Ubon Ratchathani",219,TH,Thailand,34,,15.24484530,104.84729950
3527,"Udon Thani",219,TH,Thailand,41,,17.36469690,102.81589240
3551,"Uthai Thani",219,TH,Thailand,61,,15.38350010,100.02455270
3489,Uttaradit,219,TH,Thailand,53,,17.62008860,100.09929420
3493,Yala,219,TH,Thailand,95,,44.05791170,-123.16538480
3521,Yasothon,219,TH,Thailand,35,,15.79264100,104.14528270
3601,Acklins,17,BS,"The Bahamas",AK,,22.36577080,-74.05351260
3628,"Acklins and Crooked Islands",17,BS,"The Bahamas",AC,,22.36577080,-74.05351260
3593,"Berry Islands",17,BS,"The Bahamas",BY,,25.62500420,-77.82522030
3629,Bimini,17,BS,"The Bahamas",BI,,24.64153250,-79.85062260
3605,"Black Point",17,BS,"The Bahamas",BP,,41.39510240,-71.46505560
3611,"Cat Island",17,BS,"The Bahamas",CI,,30.22801360,-89.10149330
3603,"Central Abaco",17,BS,"The Bahamas",CO,,26.35550290,-77.14851630
3631,"Central Andros",17,BS,"The Bahamas",CS,,24.46884820,-77.97386500
3596,"Central Eleuthera",17,BS,"The Bahamas",CE,,25.13620370,-76.14359150
3621,"Crooked Island",17,BS,"The Bahamas",CK,,22.63909820,-74.00650900
3614,"East Grand Bahama",17,BS,"The Bahamas",EG,,26.65828230,-78.22482910
3612,Exuma,17,BS,"The Bahamas",EX,,23.61925980,-75.96954650
3626,Freeport,17,BS,"The Bahamas",FP,,42.29668610,-89.62122710
3619,"Fresh Creek",17,BS,"The Bahamas",FC,,40.65437560,-73.89479390
3597,"Governor's Harbour",17,BS,"The Bahamas",GH,,25.19480960,-76.24396220
3632,"Grand Cay",17,BS,"The Bahamas",GC,,27.21626150,-78.32305590
3595,"Green Turtle Cay",17,BS,"The Bahamas",GT,,26.77471070,-77.32957080
3613,"Harbour Island",17,BS,"The Bahamas",HI,,25.50011000,-76.63405110
3598,"High Rock",17,BS,"The Bahamas",HR,,46.68434150,-121.90174610
3624,"Hope Town",17,BS,"The Bahamas",HT,,26.50095040,-76.99598720
3609,Inagua,17,BS,"The Bahamas",IN,,21.06560660,-73.32370800
3618,"Kemps Bay",17,BS,"The Bahamas",KB,,24.02364000,-77.54534900
3610,"Long Island",17,BS,"The Bahamas",LI,,40.78914200,-73.13496100
3625,"Mangrove Cay",17,BS,"The Bahamas",MC,,24.14814250,-77.76809520
3604,"Marsh Harbour",17,BS,"The Bahamas",MH,,26.52416530,-77.09098090
3633,"Mayaguana District",17,BS,"The Bahamas",MG,,22.40177140,-73.06413960
4881,"New Providence",17,BS,"The Bahamas",NP,,40.69843480,-74.40154050
3594,"Nichollstown and Berry Islands",17,BS,"The Bahamas",NB,,25.72362340,-77.83101040
3616,"North Abaco",17,BS,"The Bahamas",NO,,26.78716970,-77.43577390
3617,"North Andros",17,BS,"The Bahamas",NS,,24.70638050,-78.01953870
3602,"North Eleuthera",17,BS,"The Bahamas",NE,,25.46475170,-76.67592200
3615,"Ragged Island",17,BS,"The Bahamas",RI,,41.59743100,-71.26020200
3623,"Rock Sound",17,BS,"The Bahamas",RS,,39.01424430,-95.67089890
3600,"Rum Cay District",17,BS,"The Bahamas",RC,,23.68546760,-74.83901620
3620,"San Salvador and Rum Cay",17,BS,"The Bahamas",SR,,23.68546760,-74.83901620
3627,"San Salvador Island",17,BS,"The Bahamas",SS,,24.07755460,-74.47600880
3606,"Sandy Point",17,BS,"The Bahamas",SP,,39.01454640,-76.39989250
3608,"South Abaco",17,BS,"The Bahamas",SO,,26.06405910,-77.26350380
3622,"South Andros",17,BS,"The Bahamas",SA,,23.97135560,-77.60778650
3607,"South Eleuthera",17,BS,"The Bahamas",SE,,24.77085620,-76.21314740
3630,"Spanish Wells",17,BS,"The Bahamas",SW,,26.32505990,-81.79803280
3599,"West Grand Bahama",17,BS,"The Bahamas",WG,,26.65944700,-78.52065000
4520,"Aileu municipality",63,TL,Timor-Leste,AL,,-8.70439940,125.60954740
4518,"Ainaro Municipality",63,TL,Timor-Leste,AN,,-9.01131710,125.52200120
4521,"Baucau Municipality",63,TL,Timor-Leste,BA,,-8.47143080,126.45759910
4525,"Bobonaro Municipality",63,TL,Timor-Leste,BO,,-8.96554060,125.25879640
4522,"Cova Lima Municipality",63,TL,Timor-Leste,CO,,-9.26503750,125.25879640
4524,"Dili municipality",63,TL,Timor-Leste,DI,,-8.24496130,125.58766970
4516,"Ermera District",63,TL,Timor-Leste,ER,,-8.75248020,125.39872940
4523,"Lautém Municipality",63,TL,Timor-Leste,LA,,-8.36423070,126.90438450
4515,"Liquiçá Municipality",63,TL,Timor-Leste,LI,,-8.66740950,125.25879640
4517,"Manatuto District",63,TL,Timor-Leste,MT,,-8.51556080,126.01592550
4519,"Manufahi Municipality",63,TL,Timor-Leste,MF,,-9.01454950,125.82799590
4514,"Viqueque Municipality",63,TL,Timor-Leste,VI,,-8.85979180,126.36335160
2575,"Centrale Region",220,TG,Togo,C,,8.65860290,1.05861350
2579,"Kara Region",220,TG,Togo,K,,9.72163930,1.05861350
2576,Maritime,220,TG,Togo,M,,41.65514930,-83.52784670
2577,"Plateaux Region",220,TG,Togo,P,,7.61013780,1.05861350
2578,"Savanes Region",220,TG,Togo,S,,10.52917810,0.52578230
3913,Haʻapai,222,TO,Tonga,02,,-19.75000000,-174.36666700
3915,ʻEua,222,TO,Tonga,01,,37.09024000,-95.71289100
3914,Niuas,222,TO,Tonga,03,,-15.95940000,-173.78300000
3912,Tongatapu,222,TO,Tonga,04,,-21.14659680,-175.25154820
3911,Vavaʻu,222,TO,Tonga,05,,-18.62275600,-173.99029820
3362,Arima,223,TT,"Trinidad and Tobago",ARI,,46.79316040,-71.25843110
3366,Chaguanas,223,TT,"Trinidad and Tobago",CHA,,10.51683870,-61.41144820
3354,"Couva-Tabaquite-Talparo Regional Corporation",223,TT,"Trinidad and Tobago",CTT,,10.42971450,-61.37352100
3367,"Diego Martin Regional Corporation",223,TT,"Trinidad and Tobago",DMN,,10.73622860,-61.55448360
3355,"Eastern Tobago",223,TT,"Trinidad and Tobago",ETO,,11.29793480,-60.55885240
3365,"Penal-Debe Regional Corporation",223,TT,"Trinidad and Tobago",PED,,10.13374020,-61.44354740
3360,"Point Fortin",223,TT,"Trinidad and Tobago",PTF,,10.17027370,-61.67133860
3363,"Port of Spain",223,TT,"Trinidad and Tobago",POS,,10.66031960,-61.50856250
3368,"Princes Town Regional Corporation",223,TT,"Trinidad and Tobago",PRT,,10.17867460,-61.28019960
3356,"Rio Claro-Mayaro Regional Corporation",223,TT,"Trinidad and Tobago",MRC,,10.24128320,-61.09372060
3359,"San Fernando",223,TT,"Trinidad and Tobago",SFO,,34.28194610,-118.43897190
3357,"San Juan-Laventille Regional Corporation",223,TT,"Trinidad and Tobago",SJL,,10.69085780,-61.45522130
3361,"Sangre Grande Regional Corporation",223,TT,"Trinidad and Tobago",SGE,,10.58529390,-61.13158130
3364,"Siparia Regional Corporation",223,TT,"Trinidad and Tobago",SIP,,10.12456260,-61.56032440
3358,"Tunapuna-Piarco Regional Corporation",223,TT,"Trinidad and Tobago",TUP,,10.68590960,-61.30352480
3353,"Western Tobago",223,TT,"Trinidad and Tobago",WTO,,11.18970720,-60.77954520
2550,Ariana,224,TN,Tunisia,12,governorate,36.99227510,10.12551640
2572,Béja,224,TN,Tunisia,31,governorate,35.17227160,8.83076260
2566,"Ben Arous",224,TN,Tunisia,13,governorate,36.64356060,10.21515780
2551,Bizerte,224,TN,Tunisia,23,governorate,37.16093970,9.63413500
2558,Gabès,224,TN,Tunisia,81,governorate,33.94596480,9.72326730
2556,Gafsa,224,TN,Tunisia,71,governorate,34.37885050,8.66005860
2552,Jendouba,224,TN,Tunisia,32,governorate,36.71818620,8.74811670
2564,Kairouan,224,TN,Tunisia,41,governorate,35.67116630,10.10054690
2570,Kasserine,224,TN,Tunisia,42,governorate,35.08091480,8.66005860
2562,Kebili,224,TN,Tunisia,73,governorate,33.70715510,8.97146230
2561,Kef,224,TN,Tunisia,33,governorate,36.12305120,8.66005860
2568,Mahdia,224,TN,Tunisia,53,governorate,35.33525580,10.89030990
2555,Manouba,224,TN,Tunisia,14,governorate,36.84465040,9.85714160
2560,Medenine,224,TN,Tunisia,82,governorate,33.22805650,10.89030990
2553,Monastir,224,TN,Tunisia,52,governorate,35.76425150,10.81128850
5116,Nabeul,224,TN,Tunisia,21,governorate,36.45245910,10.68032220
2557,Sfax,224,TN,Tunisia,61,governorate,34.86065810,10.34978950
2567,"Sidi Bouzid",224,TN,Tunisia,43,governorate,35.03543860,9.48393920
2563,Siliana,224,TN,Tunisia,34,governorate,36.08872080,9.36453350
2571,Sousse,224,TN,Tunisia,51,governorate,35.90222670,10.34978950
2559,Tataouine,224,TN,Tunisia,83,governorate,32.13441220,10.08072980
2569,Tozeur,224,TN,Tunisia,72,governorate,33.97894910,8.04651850
2554,Tunis,224,TN,Tunisia,11,governorate,36.83749460,10.19273890
2565,Zaghouan,224,TN,Tunisia,22,governorate,36.40911880,10.14231720
2212,Adana,225,TR,Turkey,01,province,37.26123150,35.39050460
2155,Adıyaman,225,TR,Turkey,02,province,37.90782910,38.48499230
2179,Afyonkarahisar,225,TR,Turkey,03,province,38.73910990,30.71200230
2193,Ağrı,225,TR,Turkey,04,province,39.62692180,43.02159650
2210,Aksaray,225,TR,Turkey,68,province,38.33520430,33.97500180
2161,Amasya,225,TR,Turkey,05,province,40.65166080,35.90379660
2217,Ankara,225,TR,Turkey,06,province,39.78052450,32.71813750
2169,Antalya,225,TR,Turkey,07,province,37.09516720,31.07937050
2185,Ardahan,225,TR,Turkey,75,province,41.11129640,42.78316740
2191,Artvin,225,TR,Turkey,08,province,41.07866400,41.76282230
2187,Aydın,225,TR,Turkey,09,province,37.81170330,28.48639630
2175,Balıkesir,225,TR,Turkey,10,province,39.76167820,28.11226790
2148,Bartın,225,TR,Turkey,74,province,41.58105090,32.46097940
2194,Batman,225,TR,Turkey,72,province,37.83624960,41.36057390
2177,Bayburt,225,TR,Turkey,69,province,40.26032000,40.22804800
2221,Bilecik,225,TR,Turkey,11,province,40.05665550,30.06652360
2153,Bingöl,225,TR,Turkey,12,province,39.06263540,40.76960950
2215,Bitlis,225,TR,Turkey,13,province,38.65231330,42.42020280
2172,Bolu,225,TR,Turkey,14,province,40.57597660,31.57880860
2209,Burdur,225,TR,Turkey,15,province,37.46126690,30.06652360
2163,Bursa,225,TR,Turkey,16,province,40.06554590,29.23207840
2216,Çanakkale,225,TR,Turkey,17,province,40.05101040,26.98524220
2168,Çankırı,225,TR,Turkey,18,province,40.53690730,33.58838930
2173,Çorum,225,TR,Turkey,19,province,40.49982110,34.59862630
2157,Denizli,225,TR,Turkey,20,province,37.61283950,29.23207840
2226,Diyarbakır,225,TR,Turkey,21,province,38.10663720,40.54268960
2202,Düzce,225,TR,Turkey,81,province,40.87705310,31.31927130
2151,Edirne,225,TR,Turkey,22,province,41.15172220,26.51379640
2159,Elazığ,225,TR,Turkey,23,province,38.49648040,39.21990290
2160,Erzincan,225,TR,Turkey,24,province,39.76819140,39.05013060
2165,Erzurum,225,TR,Turkey,25,province,40.07467990,41.66945620
2164,Eskişehir,225,TR,Turkey,26,province,39.63296570,31.26263660
2203,Gaziantep,225,TR,Turkey,27,province,37.07638820,37.38272340
2186,Giresun,225,TR,Turkey,28,province,40.64616720,38.59355110
2204,Gümüşhane,225,TR,Turkey,29,province,40.28036730,39.31432530
2190,Hakkâri,225,TR,Turkey,30,province,37.44593190,43.74498410
2211,Hatay,225,TR,Turkey,31,province,36.40184880,36.34980970
2166,Iğdır,225,TR,Turkey,76,province,39.88798410,44.00483650
2222,Isparta,225,TR,Turkey,32,province,38.02114640,31.07937050
2170,İstanbul,225,TR,Turkey,34,province,41.16343020,28.76644080
2205,İzmir,225,TR,Turkey,35,province,38.35916930,27.26761160
2227,Kahramanmaraş,225,TR,Turkey,46,province,37.75030360,36.95410700
2223,Karabük,225,TR,Turkey,78,province,41.18748900,32.74174190
2184,Karaman,225,TR,Turkey,70,province,37.24363360,33.61757700
2208,Kars,225,TR,Turkey,36,province,40.28076360,42.99195270
2197,Kastamonu,225,TR,Turkey,37,province,41.41038630,33.69983340
2200,Kayseri,225,TR,Turkey,38,province,38.62568540,35.74068820
2154,Kilis,225,TR,Turkey,79,province,36.82047750,37.16873390
2178,Kırıkkale,225,TR,Turkey,71,province,39.88768780,33.75552480
2176,Kırklareli,225,TR,Turkey,39,province,41.72597950,27.48383900
2180,Kırşehir,225,TR,Turkey,40,province,39.22689050,33.97500180
2195,Kocaeli,225,TR,Turkey,41,province,40.85327040,29.88152030
2171,Konya,225,TR,Turkey,42,province,37.98381340,32.71813750
2149,Kütahya,225,TR,Turkey,43,province,39.35813700,29.60354950
2158,Malatya,225,TR,Turkey,44,province,38.40150570,37.95362980
2198,Manisa,225,TR,Turkey,45,province,38.84193730,28.11226790
2224,Mardin,225,TR,Turkey,47,province,37.34429290,40.61964870
2156,Mersin,225,TR,Turkey,33,province,36.81208580,34.64147500
2182,Muğla,225,TR,Turkey,48,province,37.18358190,28.48639630
2162,Muş,225,TR,Turkey,49,province,38.94618880,41.75389310
2196,Nevşehir,225,TR,Turkey,50,province,38.69393990,34.68565090
2189,Niğde,225,TR,Turkey,51,province,38.09930860,34.68565090
2174,Ordu,225,TR,Turkey,52,province,40.79905800,37.38990050
2214,Osmaniye,225,TR,Turkey,80,province,37.21302580,36.17626150
2219,Rize,225,TR,Turkey,53,province,40.95814970,40.92269850
2150,Sakarya,225,TR,Turkey,54,province,40.78885500,30.40595400
2220,Samsun,225,TR,Turkey,55,province,41.18648590,36.13226780
2183,Şanlıurfa,225,TR,Turkey,63,province,37.35691020,39.15436770
2207,Siirt,225,TR,Turkey,56,province,37.86588620,42.14945230
4854,Sinop,225,TR,Turkey,57,province,41.55947490,34.85805320
2181,Sivas,225,TR,Turkey,58,province,39.44880390,37.12944970
2225,Şırnak,225,TR,Turkey,73,province,37.41874810,42.49183380
2167,Tekirdağ,225,TR,Turkey,59,province,41.11212270,27.26761160
2199,Tokat,225,TR,Turkey,60,province,40.39027130,36.62518630
2206,Trabzon,225,TR,Turkey,61,province,40.79924100,39.58479440
2192,Tunceli,225,TR,Turkey,62,province,39.30735540,39.43877780
2201,Uşak,225,TR,Turkey,64,province,38.54313190,29.23207840
2152,Van,225,TR,Turkey,65,province,38.36794170,43.71827870
2218,Yalova,225,TR,Turkey,77,province,40.57759860,29.20883030
2188,Yozgat,225,TR,Turkey,66,province,39.72719790,35.10778580
2213,Zonguldak,225,TR,Turkey,67,province,41.31249170,31.85982510
3374,"Ahal Region",226,TM,Turkmenistan,A,,38.63993980,59.47209040
3371,Ashgabat,226,TM,Turkmenistan,S,,37.96007660,58.32606290
3372,"Balkan Region",226,TM,Turkmenistan,B,,41.81014720,21.09373110
3373,"Daşoguz Region",226,TM,Turkmenistan,D,,41.83687370,59.96519040
3370,"Lebap Region",226,TM,Turkmenistan,L,,38.12724620,64.71624150
3369,"Mary Region",226,TM,Turkmenistan,M,,36.94816230,62.45041540
3951,Funafuti,228,TV,Tuvalu,FUN,,-8.52114710,179.19619260
3947,Nanumanga,228,TV,Tuvalu,NMG,,-6.28580190,176.31992800
3949,Nanumea,228,TV,Tuvalu,NMA,,-5.68816170,176.13701480
3946,"Niutao Island Council",228,TV,Tuvalu,NIT,,-6.10642580,177.34384290
3948,Nui,228,TV,Tuvalu,NUI,,-7.23887680,177.14852320
3952,Nukufetau,228,TV,Tuvalu,NKF,,-8.00000000,178.50000000
3953,Nukulaelae,228,TV,Tuvalu,NKL,,-9.38111100,179.85222200
3950,Vaitupu,228,TV,Tuvalu,VAI,,-7.47673270,178.67476750
329,"Abim District",229,UG,Uganda,314,,2.70669800,33.65953370
361,"Adjumani District",229,UG,Uganda,301,,3.25485270,31.71954590
392,"Agago District",229,UG,Uganda,322,,2.92508200,33.34861470
344,"Alebtong District",229,UG,Uganda,323,,2.25457730,33.34861470
416,"Amolatar District",229,UG,Uganda,315,,1.60544020,32.80844960
353,"Amudat District",229,UG,Uganda,324,,1.79162240,34.90655100
352,"Amuria District",229,UG,Uganda,216,,2.03017000,33.64275330
335,"Amuru District",229,UG,Uganda,316,,2.96678780,32.08374450
328,"Apac District",229,UG,Uganda,302,,1.87302630,32.62774550
447,"Arua District",229,UG,Uganda,303,,2.99598460,31.17103890
441,"Budaka District",229,UG,Uganda,217,,1.10162770,33.93039910
349,"Bududa District",229,UG,Uganda,218,,1.00296930,34.33381230
387,"Bugiri District",229,UG,Uganda,201,,0.53161270,33.75177230
391,"Buhweju District",229,UG,Uganda,420,,-0.29113590,30.29741990
377,"Buikwe District",229,UG,Uganda,117,,0.31440460,32.98883190
343,"Bukedea District",229,UG,Uganda,219,,1.35568980,34.10867930
375,"Bukomansimbi District",229,UG,Uganda,118,,-0.14327520,31.60548930
385,"Bukwo District",229,UG,Uganda,220,,1.28186510,34.72987650
428,"Bulambuli District",229,UG,Uganda,225,,1.47988460,34.37544140
389,"Buliisa District",229,UG,Uganda,416,,2.02996070,31.53700030
419,"Bundibugyo District",229,UG,Uganda,401,,0.68517630,30.02029640
381,"Bunyangabu District",229,UG,Uganda,430,,0.48709180,30.20510960
386,"Bushenyi District",229,UG,Uganda,402,,-0.48709180,30.20510960
431,"Busia District",229,UG,Uganda,202,,0.40447310,34.01958270
365,"Butaleja District",229,UG,Uganda,221,,0.84749220,33.84112880
384,"Butambala District",229,UG,Uganda,119,,0.17425000,32.10646680
388,"Butebo District",229,UG,Uganda,233,,1.21411240,33.90808960
414,"Buvuma District",229,UG,Uganda,120,,-0.37649120,33.25879300
380,"Buyende District",229,UG,Uganda,226,,1.24136820,33.12390490
396,"Central Region",229,UG,Uganda,C,,44.29687500,-94.74017330
341,"Dokolo District",229,UG,Uganda,317,,1.96364210,33.03387670
372,"Eastern Region",229,UG,Uganda,E,,6.23740360,-0.45023680
366,"Gomba District",229,UG,Uganda,121,,0.22297910,31.67393710
413,"Gulu District",229,UG,Uganda,304,,2.81857760,32.44672380
339,"Ibanda District",229,UG,Uganda,417,,-0.09648900,30.57395790
340,"Iganga District",229,UG,Uganda,203,,0.66001370,33.48319060
383,"Isingiro District",229,UG,Uganda,418,,-0.84354300,30.80394740
367,"Jinja District",229,UG,Uganda,204,,0.53437430,33.30371430
434,"Kaabong District",229,UG,Uganda,318,,3.51262150,33.97500180
426,"Kabale District",229,UG,Uganda,404,,-1.24930840,30.06652360
326,"Kabarole District",229,UG,Uganda,405,,0.58507910,30.25127280
336,"Kaberamaido District",229,UG,Uganda,213,,1.69633220,33.21385100
403,"Kagadi District",229,UG,Uganda,427,,0.94007610,30.81256380
399,"Kakumiro District",229,UG,Uganda,428,,0.78080350,31.32413890
405,"Kalangala District",229,UG,Uganda,101,,-0.63505780,32.53727410
398,"Kaliro District",229,UG,Uganda,222,,1.04311070,33.48319060
394,"Kalungu District",229,UG,Uganda,122,,-0.09528310,31.76513620
382,"Kampala District",229,UG,Uganda,102,,0.34759640,32.58251970
334,"Kamuli District",229,UG,Uganda,205,,0.91871070,33.12390490
360,"Kamwenge District",229,UG,Uganda,413,,0.22579300,30.48184460
373,"Kanungu District",229,UG,Uganda,414,,-0.81952530,29.74260400
432,"Kapchorwa District",229,UG,Uganda,206,,1.33502050,34.39763560
440,"Kasese District",229,UG,Uganda,406,,0.06462850,30.06652360
420,"Katakwi District",229,UG,Uganda,207,,1.97310300,34.06414190
368,"Kayunga District",229,UG,Uganda,112,,0.98601820,32.85357550
436,"Kibaale District",229,UG,Uganda,407,,0.90668020,31.07937050
347,"Kiboga District",229,UG,Uganda,103,,0.96575900,31.71954590
338,"Kibuku District",229,UG,Uganda,227,,1.04528740,33.79925360
355,"Kiruhura District",229,UG,Uganda,419,,-0.19279980,30.80394740
346,"Kiryandongo District",229,UG,Uganda,421,,2.01799070,32.08374450
409,"Kisoro District",229,UG,Uganda,408,,-1.22094300,29.64991620
348,"Kitgum District",229,UG,Uganda,305,,3.33968290,33.16888830
345,"Koboko District",229,UG,Uganda,319,,3.52370580,31.03351000
401,"Kole District",229,UG,Uganda,325,,2.37010970,32.76330360
443,"Kotido District",229,UG,Uganda,306,,3.04156790,33.88577470
425,"Kumi District",229,UG,Uganda,208,,1.48769990,33.93039910
369,"Kween District",229,UG,Uganda,228,,1.44387900,34.59713200
325,"Kyankwanzi District",229,UG,Uganda,123,,1.09660370,31.71954590
437,"Kyegegwa District",229,UG,Uganda,422,,0.48181930,31.05500930
402,"Kyenjojo District",229,UG,Uganda,415,,0.60929230,30.64012310
448,"Kyotera District",229,UG,Uganda,125,,-0.63589880,31.54556370
411,"Lamwo District",229,UG,Uganda,326,,3.57075680,32.53727410
342,"Lira District",229,UG,Uganda,307,,2.23161690,32.94376670
445,"Luuka District",229,UG,Uganda,229,,0.72505990,33.30371430
433,"Luwero District",229,UG,Uganda,104,,0.82711180,32.62774550
417,"Lwengo District",229,UG,Uganda,124,,-0.41652880,31.39989950
376,"Lyantonde District",229,UG,Uganda,114,,-0.22406960,31.21684660
438,"Manafwa District",229,UG,Uganda,223,,0.90635990,34.28660910
421,"Maracha District",229,UG,Uganda,320,,3.28731270,30.94030230
356,"Masaka District",229,UG,Uganda,105,,-0.44636910,31.90179540
354,"Masindi District",229,UG,Uganda,409,,1.49203630,31.71954590
418,"Mayuge District",229,UG,Uganda,214,,-0.21829820,33.57280270
350,"Mbale District",229,UG,Uganda,209,,1.03442740,34.19768820
415,"Mbarara District",229,UG,Uganda,410,,-0.60715960,30.65450220
435,"Mitooma District",229,UG,Uganda,423,,-0.61932760,30.02029640
364,"Mityana District",229,UG,Uganda,115,,0.44548450,32.08374450
395,"Moroto District",229,UG,Uganda,308,,2.61685450,34.59713200
363,"Moyo District",229,UG,Uganda,309,,3.56964640,31.67393710
327,"Mpigi District",229,UG,Uganda,106,,0.22735280,32.32492360
371,"Mubende District",229,UG,Uganda,107,,0.57727580,31.53700030
410,"Mukono District",229,UG,Uganda,108,,0.28354760,32.76330360
393,"Nakapiripirit District",229,UG,Uganda,311,,1.96061730,34.59713200
423,"Nakaseke District",229,UG,Uganda,116,,1.22308480,32.08374450
406,"Nakasongola District",229,UG,Uganda,109,,1.34897210,32.44672380
351,"Namayingo District",229,UG,Uganda,230,,-0.28035750,33.75177230
400,"Namisindwa District",229,UG,Uganda,234,,0.90710100,34.35740370
337,"Namutumba District",229,UG,Uganda,224,,0.84926100,33.66233010
430,"Napak District",229,UG,Uganda,327,,2.36299450,34.24215970
446,"Nebbi District",229,UG,Uganda,310,,2.44093920,31.35416310
424,"Ngora District",229,UG,Uganda,231,,1.49081150,33.75177230
332,"Northern Region",229,UG,Uganda,N,,9.54392690,-0.90566230
422,"Ntoroko District",229,UG,Uganda,424,,1.07881780,30.38966510
404,"Ntungamo District",229,UG,Uganda,411,,-0.98073410,30.25127280
378,"Nwoya District",229,UG,Uganda,328,,2.56244400,31.90179540
374,"Omoro District",229,UG,Uganda,331,,2.71522300,32.49200880
390,"Otuke District",229,UG,Uganda,329,,2.52140590,33.34861470
397,"Oyam District",229,UG,Uganda,321,,2.27762810,32.44672380
408,"Pader District",229,UG,Uganda,312,,2.94306820,32.80844960
357,"Pakwach District",229,UG,Uganda,332,,2.46071410,31.49417380
412,"Pallisa District",229,UG,Uganda,210,,1.23242060,33.75177230
439,"Rakai District",229,UG,Uganda,110,,-0.70691350,31.53700030
358,"Rubanda District",229,UG,Uganda,429,,-1.18611900,29.84535760
442,"Rubirizi District",229,UG,Uganda,425,,-0.26424100,30.10840330
331,"Rukiga District",229,UG,Uganda,431,,-1.13263370,30.04341200
324,"Rukungiri District",229,UG,Uganda,412,,-0.75184900,29.92779470
427,"Sembabule District",229,UG,Uganda,111,,0.06377150,31.35416310
333,"Serere District",229,UG,Uganda,232,,1.49940330,33.54900780
407,"Sheema District",229,UG,Uganda,426,,-0.55152980,30.38966510
429,"Sironko District",229,UG,Uganda,215,,1.23022740,34.24910640
444,"Soroti District",229,UG,Uganda,211,,1.72291170,33.52800720
359,"Tororo District",229,UG,Uganda,212,,0.68709940,34.06414190
362,"Wakiso District",229,UG,Uganda,113,,0.06301900,32.44672380
370,"Western Region",229,UG,Uganda,W,,40.76672150,-111.88772030
330,"Yumbe District",229,UG,Uganda,313,,3.46980230,31.24832910
379,"Zombo District",229,UG,Uganda,330,,2.55442930,30.94173680
4689,"Autonomous Republic of Crimea",230,UA,Ukraine,43,republic,44.95211700,34.10241700
4680,"Cherkaska oblast",230,UA,Ukraine,71,region,49.44443300,32.05976700
4692,"Chernihivska oblast",230,UA,Ukraine,74,region,51.49820000,31.28934990
4678,"Chernivetska oblast",230,UA,Ukraine,77,region,48.29168300,25.93521700
4675,"Dnipropetrovska oblast",230,UA,Ukraine,12,region,48.46471700,35.04618300
4691,"Donetska oblast",230,UA,Ukraine,14,region,48.01588300,37.80285000
4682,"Ivano-Frankivska oblast",230,UA,Ukraine,26,region,48.92263300,24.71111700
4686,"Kharkivska oblast",230,UA,Ukraine,63,region,49.99350000,36.23038300
4684,"Khersonska oblast",230,UA,Ukraine,65,region,46.63541700,32.61686700
4681,"Khmelnytska oblast",230,UA,Ukraine,68,region,49.42298300,26.98713310
4677,"Kirovohradska oblast",230,UA,Ukraine,35,region,48.50793300,32.26231700
4676,Kyiv,230,UA,Ukraine,30,city,50.45010000,30.52340000
4671,"Kyivska oblast",230,UA,Ukraine,32,region,50.05295060,30.76671340
4673,"Luhanska oblast",230,UA,Ukraine,09,region,48.57404100,39.30781500
4672,"Lvivska oblast",230,UA,Ukraine,46,region,49.83968300,24.02971700
4679,"Mykolaivska oblast",230,UA,Ukraine,48,region,46.97503300,31.99458290
4688,"Odeska oblast",230,UA,Ukraine,51,region,46.48458300,30.73260000
5071,"Poltavska oblast",230,UA,Ukraine,53,region,49.64291960,32.66753390
4683,"Rivnenska oblast",230,UA,Ukraine,56,region,50.61990000,26.25161700
1912,Sevastopol,230,UA,Ukraine,40,,44.61665000,33.52536710
4685,"Sumska oblast",230,UA,Ukraine,59,region,50.90770000,34.79810000
4674,"Ternopilska oblast",230,UA,Ukraine,61,region,49.55351700,25.59476700
4669,"Vinnytska oblast",230,UA,Ukraine,05,region,49.23308300,28.46821690
4690,"Volynska oblast",230,UA,Ukraine,07,region,50.74723300,25.32538300
4670,"Zakarpatska Oblast",230,UA,Ukraine,21,region,48.62080000,22.28788300
4687,"Zaporizka oblast",230,UA,Ukraine,23,region,47.83880000,35.13956700
4668,"Zhytomyrska oblast",230,UA,Ukraine,18,region,50.25465000,28.65866690
3396,"Abu Dhabi Emirate",231,AE,"United Arab Emirates",AZ,,24.45388400,54.37734380
3395,"Ajman Emirate",231,AE,"United Arab Emirates",AJ,,25.40521650,55.51364330
3391,Dubai,231,AE,"United Arab Emirates",DU,,25.20484930,55.27078280
3393,Fujairah,231,AE,"United Arab Emirates",FU,,25.12880990,56.32648490
3394,"Ras al-Khaimah",231,AE,"United Arab Emirates",RK,,25.67413430,55.98041730
3390,"Sharjah Emirate",231,AE,"United Arab Emirates",SH,,25.07539740,55.75784030
3392,"Umm al-Quwain",231,AE,"United Arab Emirates",UQ,,25.54263240,55.54753480
2463,Aberdeen,232,GB,"United Kingdom",ABE,,57.14971700,-2.09427800
2401,Aberdeenshire,232,GB,"United Kingdom",ABD,,57.28687230,-2.38156840
2387,Angus,232,GB,"United Kingdom",ANS,,37.27578860,-95.65010330
2533,Antrim,232,GB,"United Kingdom",ANT,,54.71953380,-6.20724980
2412,"Antrim and Newtownabbey",232,GB,"United Kingdom",ANN,,54.69568870,-5.94810690
2498,Ards,232,GB,"United Kingdom",ARD,,42.13918510,-87.86149720
2523,"Ards and North Down",232,GB,"United Kingdom",AND,,54.58996450,-5.59849720
2392,"Argyll and Bute",232,GB,"United Kingdom",AGB,,56.40062140,-5.48074800
2331,"Armagh City and District Council",232,GB,"United Kingdom",ARM,,54.39325920,-6.45634010
2324,"Armagh, Banbridge and Craigavon",232,GB,"United Kingdom",ABC,,54.39325920,-6.45634010
2378,"Ascension Island",232,GB,"United Kingdom",SH-AC,,-7.94671660,-14.35591580
2363,"Ballymena Borough",232,GB,"United Kingdom",BLA,,54.86426000,-6.27910740
2361,Ballymoney,232,GB,"United Kingdom",BLY,,55.07048880,-6.51737080
2315,Banbridge,232,GB,"United Kingdom",BNB,,54.34872900,-6.27048030
2499,Barnsley,232,GB,"United Kingdom",BNS,,34.29949560,-84.98458090
2339,"Bath and North East Somerset",232,GB,"United Kingdom",BAS,,51.32501020,-2.47662410
2507,Bedford,232,GB,"United Kingdom",BDF,,32.84401700,-97.14306710
2311,"Belfast district",232,GB,"United Kingdom",BFS,,54.61703660,-5.95318610
2425,Birmingham,232,GB,"United Kingdom",BIR,,33.51858920,-86.81035670
2329,"Blackburn with Darwen",232,GB,"United Kingdom",BBD,,53.69575220,-2.46829850
2451,Blackpool,232,GB,"United Kingdom",BPL,,53.81750530,-3.03567480
2530,"Blaenau Gwent County Borough",232,GB,"United Kingdom",BGW,,51.78757790,-3.20439310
2504,Bolton,232,GB,"United Kingdom",BOL,,44.37264760,-72.87876250
2342,Bournemouth,232,GB,"United Kingdom",BMH,,50.71916400,-1.88076900
2470,"Bracknell Forest",232,GB,"United Kingdom",BRC,,51.41538280,-0.75364950
2529,Bradford,232,GB,"United Kingdom",BRD,,53.79598400,-1.75939800
2452,"Bridgend County Borough",232,GB,"United Kingdom",BGE,,51.50831990,-3.58120750
2395,"Brighton and Hove",232,GB,"United Kingdom",BNH,,50.82262880,-0.13704700
2405,Buckinghamshire,232,GB,"United Kingdom",BKM,,51.80722040,-0.81276640
2459,Bury,232,GB,"United Kingdom",BUR,,53.59334980,-2.29660540
2298,"Caerphilly County Borough",232,GB,"United Kingdom",CAY,,51.66044650,-3.21787240
2517,Calderdale,232,GB,"United Kingdom",CLD,,53.72478450,-1.86583570
2423,Cambridgeshire,232,GB,"United Kingdom",CAM,,52.20529730,0.12181950
2484,Carmarthenshire,232,GB,"United Kingdom",CMN,,51.85723090,-4.31159590
2439,"Carrickfergus Borough Council",232,GB,"United Kingdom",CKF,,54.72568430,-5.80937190
2525,Castlereagh,232,GB,"United Kingdom",CSR,,54.57567900,-5.88840280
2316,"Causeway Coast and Glens",232,GB,"United Kingdom",CCG,,55.04318300,-6.67412880
2303,"Central Bedfordshire",232,GB,"United Kingdom",CBF,,52.00297440,-0.46513890
2509,Ceredigion,232,GB,"United Kingdom",CGN,,52.21914290,-3.93212560
2444,"Cheshire East",232,GB,"United Kingdom",CHE,,53.16104460,-2.21859320
2442,"Cheshire West and Chester",232,GB,"United Kingdom",CHW,,53.23029740,-2.71511170
2528,"City and County of Cardiff",232,GB,"United Kingdom",CRF,,51.48158100,-3.17909000
2433,"City and County of Swansea",232,GB,"United Kingdom",SWA,,51.62144000,-3.94364600
2413,"City of Bristol",232,GB,"United Kingdom",BST,,41.67352200,-72.94653750
2485,"City of Derby",232,GB,"United Kingdom",DER,,37.54837550,-97.24851910
2475,"City of Kingston upon Hull",232,GB,"United Kingdom",KHL,,53.76762360,-0.32741980
2318,"City of Leicester",232,GB,"United Kingdom",LCE,,52.63687780,-1.13975920
2424,"City of London",232,GB,"United Kingdom",LND,,51.51234430,-0.09098520
2359,"City of Nottingham",232,GB,"United Kingdom",NGM,,52.95478320,-1.15810860
2297,"City of Peterborough",232,GB,"United Kingdom",PTE,,44.30936360,-78.32015300
2514,"City of Plymouth",232,GB,"United Kingdom",PLY,,42.37089410,-83.46971410
2305,"City of Portsmouth",232,GB,"United Kingdom",POR,,36.83291500,-76.29755490
2294,"City of Southampton",232,GB,"United Kingdom",STH,,50.90970040,-1.40435090
2506,"City of Stoke-on-Trent",232,GB,"United Kingdom",STE,,53.00266800,-2.17940400
2372,"City of Sunderland",232,GB,"United Kingdom",SND,,54.88614890,-1.47857970
2357,"City of Westminster",232,GB,"United Kingdom",WSM,,39.57659770,-76.99721260
2489,"City of Wolverhampton",232,GB,"United Kingdom",WLV,,52.58891200,-2.15646300
2426,"City of York",232,GB,"United Kingdom",YOR,,53.95996510,-1.08729790
2450,Clackmannanshire,232,GB,"United Kingdom",CLK,,56.10753510,-3.75294090
2461,"Coleraine Borough Council",232,GB,"United Kingdom",CLR,,55.14515700,-6.67598140
2352,"Conwy County Borough",232,GB,"United Kingdom",CWY,,53.29350130,-3.72651610
2445,"Cookstown District Council",232,GB,"United Kingdom",CKT,,54.64181580,-6.74438950
2312,Cornwall,232,GB,"United Kingdom",CON,,50.26604710,-5.05271250
2406,"County Durham",232,GB,"United Kingdom",DUR,,54.72940990,-1.88115980
2438,Coventry,232,GB,"United Kingdom",COV,,52.40682200,-1.51969300
2449,"Craigavon Borough Council",232,GB,"United Kingdom",CGV,,54.39325920,-6.45634010
2334,Cumbria,232,GB,"United Kingdom",CMA,,54.57723230,-2.79748350
2389,Darlington,232,GB,"United Kingdom",DAL,,34.29987620,-79.87617410
2497,Denbighshire,232,GB,"United Kingdom",DEN,,53.18422880,-3.42249850
2403,Derbyshire,232,GB,"United Kingdom",DBY,,53.10467820,-1.56238850
2446,"Derry City and Strabane",232,GB,"United Kingdom",DRS,,55.00474430,-7.32092220
2417,"Derry City Council",232,GB,"United Kingdom",DRY,,54.96907780,-7.19583510
2491,Devon,232,GB,"United Kingdom",DEV,,50.71555910,-3.53087500
2364,Doncaster,232,GB,"United Kingdom",DNC,,53.52282000,-1.12846200
2345,Dorset,232,GB,"United Kingdom",DOR,,50.74876350,-2.34447860
2304,"Down District Council",232,GB,"United Kingdom",DOW,,54.24342870,-5.95779590
2457,Dudley,232,GB,"United Kingdom",DUD,,42.04336610,-71.92760330
2415,"Dumfries and Galloway",232,GB,"United Kingdom",DGY,,55.07010730,-3.60525810
2511,Dundee,232,GB,"United Kingdom",DND,,56.46201800,-2.97072100
2508,"Dungannon and South Tyrone Borough Council",232,GB,"United Kingdom",DGN,,54.50826840,-6.76658910
2374,"East Ayrshire",232,GB,"United Kingdom",EAY,,55.45184960,-4.26444780
2454,"East Dunbartonshire",232,GB,"United Kingdom",EDU,,55.97431620,-4.20229800
2462,"East Lothian",232,GB,"United Kingdom",ELN,,55.94933830,-2.77044640
2333,"East Renfrewshire",232,GB,"United Kingdom",ERW,,55.77047350,-4.33598210
2370,"East Riding of Yorkshire",232,GB,"United Kingdom",ERY,,53.84161680,-0.43441060
2414,"East Sussex",232,GB,"United Kingdom",ESX,,50.90859550,0.24941660
2428,Edinburgh,232,GB,"United Kingdom",EDH,,55.95325200,-3.18826700
2336,England,232,GB,"United Kingdom",ENG,,52.35551770,-1.17431970
2410,Essex,232,GB,"United Kingdom",ESS,,51.57424470,0.48567810
2344,Falkirk,232,GB,"United Kingdom",FAL,,56.00187750,-3.78391310
2366,"Fermanagh and Omagh",232,GB,"United Kingdom",FMO,,54.45135240,-7.71250180
2531,"Fermanagh District Council",232,GB,"United Kingdom",FER,,54.34479780,-7.63842180
2479,Fife,232,GB,"United Kingdom",FIF,,56.20820780,-3.14951750
2437,Flintshire,232,GB,"United Kingdom",FLN,,53.16686580,-3.14189080
2431,Gateshead,232,GB,"United Kingdom",GAT,,54.95268000,-1.60341100
2404,Glasgow,232,GB,"United Kingdom",GLG,,55.86423700,-4.25180600
2373,Gloucestershire,232,GB,"United Kingdom",GLS,,51.86421120,-2.23803350
2379,Gwynedd,232,GB,"United Kingdom",GWN,,52.92772660,-4.13348360
2466,Halton,232,GB,"United Kingdom",HAL,,43.53253720,-79.87448360
2435,Hampshire,232,GB,"United Kingdom",HAM,,51.05769480,-1.30806290
2309,Hartlepool,232,GB,"United Kingdom",HPL,,54.69174500,-1.21292600
2500,Herefordshire,232,GB,"United Kingdom",HEF,,52.07651640,-2.65441820
2369,Hertfordshire,232,GB,"United Kingdom",HRT,,51.80978230,-0.23767440
2383,Highland,232,GB,"United Kingdom",HLD,,36.29675080,-95.83803660
2388,Inverclyde,232,GB,"United Kingdom",IVC,,55.93165690,-4.68001580
2289,"Isle of Wight",232,GB,"United Kingdom",IOW,,50.69384790,-1.30473400
2343,"Isles of Scilly",232,GB,"United Kingdom",IOS,,49.92772610,-6.32749660
2464,Kent,232,GB,"United Kingdom",KEN,,41.15366740,-81.35788590
2371,Kirklees,232,GB,"United Kingdom",KIR,,53.59334320,-1.80095090
2330,Knowsley,232,GB,"United Kingdom",KWL,,53.45459400,-2.85290700
2495,Lancashire,232,GB,"United Kingdom",LAN,,53.76322540,-2.70440520
2515,"Larne Borough Council",232,GB,"United Kingdom",LRN,,54.85780030,-5.82362240
2503,Leeds,232,GB,"United Kingdom",LDS,,53.80075540,-1.54907740
2516,Leicestershire,232,GB,"United Kingdom",LEC,,52.77257100,-1.20521260
2382,"Limavady Borough Council",232,GB,"United Kingdom",LMV,,55.05168200,-6.94919440
2355,Lincolnshire,232,GB,"United Kingdom",LIN,,52.94518890,-0.16012460
2460,"Lisburn and Castlereagh",232,GB,"United Kingdom",LBC,,54.49815840,-6.13067910
2494,"Lisburn City Council",232,GB,"United Kingdom",LSB,,54.49815840,-6.13067910
2340,Liverpool,232,GB,"United Kingdom",LIV,,32.65649810,-115.47632410
2356,"London Borough of Barking and Dagenham",232,GB,"United Kingdom",BDG,,51.55406660,0.13401700
2520,"London Borough of Barnet",232,GB,"United Kingdom",BNE,,51.60496730,-0.20762950
2307,"London Borough of Bexley",232,GB,"United Kingdom",BEX,,51.45190210,0.11717860
2291,"London Borough of Brent",232,GB,"United Kingdom",BEN,,51.56728080,-0.27105680
2490,"London Borough of Bromley",232,GB,"United Kingdom",BRY,,51.36797050,0.07006200
2349,"London Borough of Camden",232,GB,"United Kingdom",CMD,,51.54547360,-0.16279020
2512,"London Borough of Croydon",232,GB,"United Kingdom",CRY,,51.38274460,-0.09851630
2532,"London Borough of Ealing",232,GB,"United Kingdom",EAL,,51.52503660,-0.34139650
2476,"London Borough of Enfield",232,GB,"United Kingdom",ENF,,51.66229090,-0.11806510
2411,"London Borough of Hackney",232,GB,"United Kingdom",HCK,,51.57344500,-0.07243760
2448,"London Borough of Hammersmith and Fulham",232,GB,"United Kingdom",HMF,,51.49901560,-0.22915000
2306,"London Borough of Haringey",232,GB,"United Kingdom",HRY,,51.59061130,-0.11097090
2385,"London Borough of Harrow",232,GB,"United Kingdom",HRW,,51.58816270,-0.34228510
2347,"London Borough of Havering",232,GB,"United Kingdom",HAV,,51.57792400,0.21208290
2376,"London Borough of Hillingdon",232,GB,"United Kingdom",HIL,,51.53518320,-0.44813780
2380,"London Borough of Hounslow",232,GB,"United Kingdom",HNS,,51.48283580,-0.38820620
2319,"London Borough of Islington",232,GB,"United Kingdom",ISL,,51.54650630,-0.10580580
2396,"London Borough of Lambeth",232,GB,"United Kingdom",LBH,,51.45714770,-0.12306810
2358,"London Borough of Lewisham",232,GB,"United Kingdom",LEW,,51.44145790,-0.01170060
2483,"London Borough of Merton",232,GB,"United Kingdom",MRT,,51.40977420,-0.21080840
2418,"London Borough of Newham",232,GB,"United Kingdom",NWM,,51.52551620,0.03521630
2397,"London Borough of Redbridge",232,GB,"United Kingdom",RDB,,51.58861210,0.08239820
2501,"London Borough of Richmond upon Thames",232,GB,"United Kingdom",RIC,,51.46130540,-0.30377090
2432,"London Borough of Southwark",232,GB,"United Kingdom",SWK,,51.48805720,-0.07628380
2313,"London Borough of Sutton",232,GB,"United Kingdom",STN,,51.35737620,-0.17527960
2390,"London Borough of Tower Hamlets",232,GB,"United Kingdom",TWH,,51.52026070,-0.02933960
2326,"London Borough of Waltham Forest",232,GB,"United Kingdom",WFT,,51.58863830,-0.01176250
2434,"London Borough of Wandsworth",232,GB,"United Kingdom",WND,,51.45682740,-0.18966380
2322,"Magherafelt District Council",232,GB,"United Kingdom",MFT,,54.75532790,-6.60774870
2398,Manchester,232,GB,"United Kingdom",MAN,,53.48075930,-2.24263050
2381,Medway,232,GB,"United Kingdom",MDW,,42.14176410,-71.39672560
2328,"Merthyr Tydfil County Borough",232,GB,"United Kingdom",MTY,,51.74674740,-3.38132750
2320,"Metropolitan Borough of Wigan",232,GB,"United Kingdom",WGN,,53.51348120,-2.61069990
2429,"Mid and East Antrim",232,GB,"United Kingdom",MEA,,54.93993410,-6.11374230
2399,"Mid Ulster",232,GB,"United Kingdom",MUL,,54.64113010,-6.75225490
2332,Middlesbrough,232,GB,"United Kingdom",MDB,,54.57422700,-1.23495600
2519,Midlothian,232,GB,"United Kingdom",MLN,,32.47533500,-97.01031810
2416,"Milton Keynes",232,GB,"United Kingdom",MIK,,52.08520380,-0.73331330
2402,Monmouthshire,232,GB,"United Kingdom",MON,,51.81161000,-2.71634170
2360,Moray,232,GB,"United Kingdom",MRY,,57.64984760,-3.31680390
2348,"Moyle District Council",232,GB,"United Kingdom",MYL,,55.20473270,-6.25317400
2351,"Neath Port Talbot County Borough",232,GB,"United Kingdom",NTL,,51.59785190,-3.78396680
2458,"Newcastle upon Tyne",232,GB,"United Kingdom",NET,,54.97825200,-1.61778000
2524,Newport,232,GB,"United Kingdom",NWP,,37.52782340,-94.10438760
2350,"Newry and Mourne District Council",232,GB,"United Kingdom",NYM,,54.17425050,-6.33919920
2534,"Newry, Mourne and Down",232,GB,"United Kingdom",NMD,,54.24342870,-5.95779590
2317,"Newtownabbey Borough Council",232,GB,"United Kingdom",NTA,,54.67924220,-5.95911020
2473,Norfolk,232,GB,"United Kingdom",NFK,,36.85076890,-76.28587260
2535,"North Ayrshire",232,GB,"United Kingdom",NAY,,55.64167310,-4.75946000
2513,"North Down Borough Council",232,GB,"United Kingdom",NDN,,54.65362970,-5.67249250
2384,"North East Lincolnshire",232,GB,"United Kingdom",NEL,,53.56682010,-0.08150660
2487,"North Lanarkshire",232,GB,"United Kingdom",NLK,,55.86624320,-3.96131440
2453,"North Lincolnshire",232,GB,"United Kingdom",NLN,,53.60555920,-0.55965820
2430,"North Somerset",232,GB,"United Kingdom",NSM,,51.38790280,-2.77810910
2521,"North Tyneside",232,GB,"United Kingdom",NTY,,55.01823990,-1.48584360
2522,"North Yorkshire",232,GB,"United Kingdom",NYK,,53.99150280,-1.54120150
2480,Northamptonshire,232,GB,"United Kingdom",NTH,,52.27299440,-0.87555150
2337,"Northern Ireland",232,GB,"United Kingdom",NIR,,54.78771490,-6.49231450
2365,Northumberland,232,GB,"United Kingdom",NBL,,55.20825420,-2.07841380
2456,Nottinghamshire,232,GB,"United Kingdom",NTT,,53.10031900,-0.99363060
2477,Oldham,232,GB,"United Kingdom",OLD,,42.20405980,-71.20481190
2314,"Omagh District Council",232,GB,"United Kingdom",OMH,,54.45135240,-7.71250180
2474,"Orkney Islands",232,GB,"United Kingdom",ORK,,58.98094010,-2.96052060
2353,"Outer Hebrides",232,GB,"United Kingdom",ELS,,57.75989180,-7.01940340
2321,Oxfordshire,232,GB,"United Kingdom",OXF,,51.76120560,-1.24646740
2486,Pembrokeshire,232,GB,"United Kingdom",PEM,,51.67407800,-4.90887850
2325,"Perth and Kinross",232,GB,"United Kingdom",PKN,,56.39538170,-3.42835470
2302,Poole,232,GB,"United Kingdom",POL,,50.71505000,-1.98724800
2441,Powys,232,GB,"United Kingdom",POW,,52.64642490,-3.32609040
2455,Reading,232,GB,"United Kingdom",RDG,,36.14866590,-95.98400120
2527,"Redcar and Cleveland",232,GB,"United Kingdom",RCC,,54.59713440,-1.07759970
2443,Renfrewshire,232,GB,"United Kingdom",RFW,,55.84665400,-4.53312590
2301,"Rhondda Cynon Taf",232,GB,"United Kingdom",RCT,,51.64902070,-3.42886920
2327,Rochdale,232,GB,"United Kingdom",RCH,,53.60971360,-2.15610000
2308,Rotherham,232,GB,"United Kingdom",ROT,,53.43260350,-1.36350090
2492,"Royal Borough of Greenwich",232,GB,"United Kingdom",GRE,,51.48346270,0.05862020
2368,"Royal Borough of Kensington and Chelsea",232,GB,"United Kingdom",KEC,,51.49908050,-0.19382530
2481,"Royal Borough of Kingston upon Thames",232,GB,"United Kingdom",KTT,,51.37811700,-0.29270900
2472,Rutland,232,GB,"United Kingdom",RUT,,43.61062370,-72.97260650
2502,"Saint Helena",232,GB,"United Kingdom",SH-HL,,-15.96501040,-5.70892410
2493,Salford,232,GB,"United Kingdom",SLF,,53.48752350,-2.29012640
2341,Sandwell,232,GB,"United Kingdom",SAW,,52.53616740,-2.01079300
2335,Scotland,232,GB,"United Kingdom",SCT,,56.49067120,-4.20264580
2346,"Scottish Borders",232,GB,"United Kingdom",SCB,,55.54856970,-2.78613880
2518,Sefton,232,GB,"United Kingdom",SFT,,53.50344490,-2.97035900
2295,Sheffield,232,GB,"United Kingdom",SHF,,36.09507430,-80.27884660
2300,"Shetland Islands",232,GB,"United Kingdom",ZET,,60.52965070,-1.26594090
2407,Shropshire,232,GB,"United Kingdom",SHR,,52.70636570,-2.74178490
2427,Slough,232,GB,"United Kingdom",SLG,,51.51053840,-0.59504060
2469,Solihull,232,GB,"United Kingdom",SOL,,52.41181100,-1.77761000
2386,Somerset,232,GB,"United Kingdom",SOM,,51.10509700,-2.92623070
2377,"South Ayrshire",232,GB,"United Kingdom",SAY,,55.45889880,-4.62919940
2400,"South Gloucestershire",232,GB,"United Kingdom",SGC,,51.52643610,-2.47284870
2362,"South Lanarkshire",232,GB,"United Kingdom",SLK,,55.67359090,-3.78196610
2409,"South Tyneside",232,GB,"United Kingdom",STY,,54.96366930,-1.44186340
2323,Southend-on-Sea,232,GB,"United Kingdom",SOS,,51.54592690,0.70771230
2290,"St Helens",232,GB,"United Kingdom",SHN,,45.85896100,-122.82123560
2447,Staffordshire,232,GB,"United Kingdom",STS,,52.87927450,-2.05718680
2488,Stirling,232,GB,"United Kingdom",STG,,56.11652270,-3.93690290
2394,Stockport,232,GB,"United Kingdom",SKP,,53.41063160,-2.15753320
2421,Stockton-on-Tees,232,GB,"United Kingdom",STT,,54.57045510,-1.32898210
2393,"Strabane District Council",232,GB,"United Kingdom",STB,,54.82738650,-7.46331030
2467,Suffolk,232,GB,"United Kingdom",SFK,,52.18724720,0.97078010
2526,Surrey,232,GB,"United Kingdom",SRY,,51.31475930,-0.55995010
2422,Swindon,232,GB,"United Kingdom",SWD,,51.55577390,-1.77971760
2367,Tameside,232,GB,"United Kingdom",TAM,,53.48058280,-2.08098910
2310,"Telford and Wrekin",232,GB,"United Kingdom",TFW,,52.74099160,-2.48685860
2468,Thurrock,232,GB,"United Kingdom",THR,,51.49345570,0.35291970
2478,Torbay,232,GB,"United Kingdom",TOB,,50.43923290,-3.53698990
2496,Torfaen,232,GB,"United Kingdom",TOF,,51.70022530,-3.04460150
2293,Trafford,232,GB,"United Kingdom",TRF,,40.38562460,-79.75893470
2375,"United Kingdom",232,GB,"United Kingdom",UKM,,55.37805100,-3.43597300
2299,"Vale of Glamorgan",232,GB,"United Kingdom",VGL,,51.40959580,-3.48481670
2465,Wakefield,232,GB,"United Kingdom",WKF,,42.50393950,-71.07233910
2338,Wales,232,GB,"United Kingdom",WLS,,52.13066070,-3.78371170
2292,Walsall,232,GB,"United Kingdom",WLL,,52.58621400,-1.98291900
2420,Warrington,232,GB,"United Kingdom",WRT,,40.24927410,-75.13406040
2505,Warwickshire,232,GB,"United Kingdom",WAR,,52.26713530,-1.46752160
2471,"West Berkshire",232,GB,"United Kingdom",WBK,,51.43082550,-1.14449270
2440,"West Dunbartonshire",232,GB,"United Kingdom",WDU,,55.94509250,-4.56462590
2354,"West Lothian",232,GB,"United Kingdom",WLN,,55.90701980,-3.55171670
2296,"West Sussex",232,GB,"United Kingdom",WSX,,50.92801430,-0.46170750
2391,Wiltshire,232,GB,"United Kingdom",WIL,,51.34919960,-1.99271050
2482,"Windsor and Maidenhead",232,GB,"United Kingdom",WNM,,51.47997120,-0.62425650
2408,Wirral,232,GB,"United Kingdom",WRL,,53.37271810,-3.07375400
2419,Wokingham,232,GB,"United Kingdom",WOK,,51.41045700,-0.83386100
2510,Worcestershire,232,GB,"United Kingdom",WOR,,52.25452250,-2.26683820
2436,"Wrexham County Borough",232,GB,"United Kingdom",WRX,,53.03013780,-3.02614870
1456,Alabama,233,US,"United States",AL,state,32.31823140,-86.90229800
1400,Alaska,233,US,"United States",AK,state,64.20084130,-149.49367330
1424,"American Samoa",233,US,"United States",AS,"outlying area",-14.27097200,-170.13221700
1434,Arizona,233,US,"United States",AZ,state,34.04892810,-111.09373110
1444,Arkansas,233,US,"United States",AR,state,35.20105000,-91.83183340
1402,"Baker Island",233,US,"United States",UM-81,"islands / groups of islands",0.19362660,-176.47690800
1416,California,233,US,"United States",CA,state,36.77826100,-119.41793240
1450,Colorado,233,US,"United States",CO,state,39.55005070,-105.78206740
1435,Connecticut,233,US,"United States",CT,state,41.60322070,-73.08774900
1399,Delaware,233,US,"United States",DE,state,38.91083250,-75.52766990
1437,"District of Columbia",233,US,"United States",DC,district,38.90719230,-77.03687070
1436,Florida,233,US,"United States",FL,state,27.66482740,-81.51575350
1455,Georgia,233,US,"United States",GA,state,32.16562210,-82.90007510
1412,Guam,233,US,"United States",GU,"outlying area",13.44430400,144.79373100
1411,Hawaii,233,US,"United States",HI,state,19.89676620,-155.58278180
1398,"Howland Island",233,US,"United States",UM-84,"islands / groups of islands",0.81132190,-176.61827360
1460,Idaho,233,US,"United States",ID,state,44.06820190,-114.74204080
1425,Illinois,233,US,"United States",IL,state,40.63312490,-89.39852830
1440,Indiana,233,US,"United States",IN,state,40.26719410,-86.13490190
1459,Iowa,233,US,"United States",IA,state,41.87800250,-93.09770200
1410,"Jarvis Island",233,US,"United States",UM-86,"islands / groups of islands",-0.37435030,-159.99672060
1428,"Johnston Atoll",233,US,"United States",UM-67,"islands / groups of islands",16.72950350,-169.53364770
1406,Kansas,233,US,"United States",KS,state,39.01190200,-98.48424650
1419,Kentucky,233,US,"United States",KY,state,37.83933320,-84.27001790
1403,"Kingman Reef",233,US,"United States",UM-89,"islands / groups of islands",6.38333300,-162.41666700
1457,Louisiana,233,US,"United States",LA,state,30.98429770,-91.96233270
1453,Maine,233,US,"United States",ME,state,45.25378300,-69.44546890
1401,Maryland,233,US,"United States",MD,state,39.04575490,-76.64127120
1433,Massachusetts,233,US,"United States",MA,state,42.40721070,-71.38243740
1426,Michigan,233,US,"United States",MI,state,44.31484430,-85.60236430
1438,"Midway Atoll",233,US,"United States",UM-71,"islands / groups of islands",28.20721680,-177.37349260
1420,Minnesota,233,US,"United States",MN,state,46.72955300,-94.68589980
1430,Mississippi,233,US,"United States",MS,state,32.35466790,-89.39852830
1451,Missouri,233,US,"United States",MO,state,37.96425290,-91.83183340
1446,Montana,233,US,"United States",MT,state,46.87968220,-110.36256580
1439,"Navassa Island",233,US,"United States",UM-76,"islands / groups of islands",18.41006890,-75.01146120
1408,Nebraska,233,US,"United States",NE,state,41.49253740,-99.90181310
1458,Nevada,233,US,"United States",NV,state,38.80260970,-116.41938900
1404,"New Hampshire",233,US,"United States",NH,state,43.19385160,-71.57239530
1417,"New Jersey",233,US,"United States",NJ,state,40.05832380,-74.40566120
1423,"New Mexico",233,US,"United States",NM,state,34.51994020,-105.87009010
1452,"New York",233,US,"United States",NY,state,40.71277530,-74.00597280
1447,"North Carolina",233,US,"United States",NC,state,35.75957310,-79.01929970
1418,"North Dakota",233,US,"United States",ND,state,47.55149260,-101.00201190
1431,"Northern Mariana Islands",233,US,"United States",MP,"outlying area",15.09790000,145.67390000
4851,Ohio,233,US,"United States",OH,state,40.41728710,-82.90712300
1421,Oklahoma,233,US,"United States",OK,state,35.46756020,-97.51642760
1415,Oregon,233,US,"United States",OR,state,43.80413340,-120.55420120
1448,"Palmyra Atoll",233,US,"United States",UM-95,"islands / groups of islands",5.88850260,-162.07866560
1422,Pennsylvania,233,US,"United States",PA,state,41.20332160,-77.19452470
1449,"Puerto Rico",233,US,"United States",PR,"outlying area",18.22083300,-66.59014900
1461,"Rhode Island",233,US,"United States",RI,state,41.58009450,-71.47742910
1443,"South Carolina",233,US,"United States",SC,state,33.83608100,-81.16372450
1445,"South Dakota",233,US,"United States",SD,state,43.96951480,-99.90181310
1454,Tennessee,233,US,"United States",TN,state,35.51749130,-86.58044730
1407,Texas,233,US,"United States",TX,state,31.96859880,-99.90181310
1432,"United States Minor Outlying Islands",233,US,"United States",UM,"outlying area",19.28231920,166.64704700
1413,"United States Virgin Islands",233,US,"United States",VI,"outlying area",18.33576500,-64.89633500
1414,Utah,233,US,"United States",UT,state,39.32098010,-111.09373110
1409,Vermont,233,US,"United States",VT,state,44.55880280,-72.57784150
1427,Virginia,233,US,"United States",VA,state,37.43157340,-78.65689420
1405,"Wake Island",233,US,"United States",UM-79,"islands / groups of islands",19.27961900,166.64993480
1462,Washington,233,US,"United States",WA,state,47.75107410,-120.74013850
1429,"West Virginia",233,US,"United States",WV,state,38.59762620,-80.45490260
1441,Wisconsin,233,US,"United States",WI,state,43.78443970,-88.78786780
1442,Wyoming,233,US,"United States",WY,state,43.07596780,-107.29028390
5212,"Baker Island",234,UM,"United States Minor Outlying Islands",81,island,0.19362660,-176.47690800
5213,"Howland Island",234,UM,"United States Minor Outlying Islands",84,island,0.81132190,-176.61827360
5214,"Jarvis Island",234,UM,"United States Minor Outlying Islands",86,island,-0.37435030,-159.99672060
5215,"Johnston Atoll",234,UM,"United States Minor Outlying Islands",67,island,16.72950350,-169.53364770
5216,"Kingman Reef",234,UM,"United States Minor Outlying Islands",89,island,6.38333300,-162.41666700
5217,"Midway Islands",234,UM,"United States Minor Outlying Islands",71,island,28.20721680,-177.37349260
5218,"Navassa Island",234,UM,"United States Minor Outlying Islands",76,island,18.41006890,-75.01146120
5219,"Palmyra Atoll",234,UM,"United States Minor Outlying Islands",95,island,5.88850260,-162.07866560
5220,"Wake Island",234,UM,"United States Minor Outlying Islands",79,island,19.27961900,166.64993480
3205,Artigas,235,UY,Uruguay,AR,,-30.61751120,-56.95945590
3213,Canelones,235,UY,Uruguay,CA,,-34.54087170,-55.93076000
3211,"Cerro Largo",235,UY,Uruguay,CL,,-32.44110320,-54.35217530
3208,Colonia,235,UY,Uruguay,CO,,-34.12946780,-57.66051840
3209,Durazno,235,UY,Uruguay,DU,,-33.02324540,-56.02846440
3203,Flores,235,UY,Uruguay,FS,,-33.57337530,-56.89450280
3217,Florida,235,UY,Uruguay,FD,,28.03594950,-82.45792890
3215,Lavalleja,235,UY,Uruguay,LA,,-33.92261750,-54.97657940
3206,Maldonado,235,UY,Uruguay,MA,,-34.55979320,-54.86285520
3218,Montevideo,235,UY,Uruguay,MO,,-34.81815870,-56.21382560
3212,Paysandú,235,UY,Uruguay,PA,,-32.06673660,-57.33647890
3210,"Río Negro",235,UY,Uruguay,RN,,-32.76763560,-57.42952070
3207,Rivera,235,UY,Uruguay,RV,,-31.48174210,-55.24357590
3216,Rocha,235,UY,Uruguay,RO,,-33.96900810,-54.02148500
3220,Salto,235,UY,Uruguay,SA,,-31.38802800,-57.96124550
3204,"San José",235,UY,Uruguay,SJ,,37.34929680,-121.90560490
3219,Soriano,235,UY,Uruguay,SO,,-33.51027920,-57.74981030
3221,Tacuarembó,235,UY,Uruguay,TA,,-31.72068370,-55.98598870
3214,"Treinta y Tres",235,UY,Uruguay,TT,,-33.06850860,-54.28586270
2540,"Andijan Region",236,UZ,Uzbekistan,AN,,40.76859410,72.23637900
2541,"Bukhara Region",236,UZ,Uzbekistan,BU,,40.25041620,63.20321510
2538,"Fergana Region",236,UZ,Uzbekistan,FA,,40.45680810,71.28742090
2545,"Jizzakh Region",236,UZ,Uzbekistan,JI,,40.47064150,67.57085360
2548,Karakalpakstan,236,UZ,Uzbekistan,QR,,43.80413340,59.44579880
2537,"Namangan Region",236,UZ,Uzbekistan,NG,,41.05100370,71.09731700
2542,"Navoiy Region",236,UZ,Uzbekistan,NW,,42.69885750,64.63376850
2543,"Qashqadaryo Region",236,UZ,Uzbekistan,QA,,38.89862310,66.04635340
2544,"Samarqand Region",236,UZ,Uzbekistan,SA,,39.62701200,66.97497310
2547,"Sirdaryo Region",236,UZ,Uzbekistan,SI,,40.38638080,68.71549750
2546,"Surxondaryo Region",236,UZ,Uzbekistan,SU,,37.94090050,67.57085360
2536,Tashkent,236,UZ,Uzbekistan,TK,,41.29949580,69.24007340
2549,"Tashkent Region",236,UZ,Uzbekistan,TO,,41.22132340,69.85974060
2539,"Xorazm Region",236,UZ,Uzbekistan,XO,,41.35653360,60.85666860
4775,Malampa,237,VU,Vanuatu,MAP,,-16.40114050,167.60778650
4773,Penama,237,VU,Vanuatu,PAM,,-15.37957580,167.90531820
4776,Sanma,237,VU,Vanuatu,SAM,,-15.48400170,166.91820970
4774,Shefa,237,VU,Vanuatu,SEE,,32.80576500,35.16997100
4777,Tafea,237,VU,Vanuatu,TAE,,-18.72378270,169.06450560
4772,Torba,237,VU,Vanuatu,TOB,,37.07653000,27.45657300
2044,Amazonas,239,VE,Venezuela,Z,state,-3.41684270,-65.85606460
2050,Anzoátegui,239,VE,Venezuela,B,state,8.59130730,-63.95861110
4856,Apure,239,VE,Venezuela,C,state,6.92694830,-68.52471490
2047,Aragua,239,VE,Venezuela,D,state,10.06357580,-67.28478750
2049,Barinas,239,VE,Venezuela,E,state,8.62314980,-70.23710450
2039,Bolívar,239,VE,Venezuela,F,state,37.61448380,-93.41047490
2040,Carabobo,239,VE,Venezuela,G,state,10.11764330,-68.04775090
2034,Cojedes,239,VE,Venezuela,H,state,9.38166820,-68.33392750
2051,"Delta Amacuro",239,VE,Venezuela,Y,state,8.84993070,-61.14031960
4855,"Distrito Capital",239,VE,Venezuela,A,"capital district",41.26148460,-95.93108070
2035,Falcón,239,VE,Venezuela,I,state,11.18106740,-69.85974060
2046,"Federal Dependencies of Venezuela",239,VE,Venezuela,W,"federal dependency",10.93770530,-65.35695730
2045,Guárico,239,VE,Venezuela,J,state,8.74893090,-66.23671720
2055,"La Guaira",239,VE,Venezuela,X,state,29.30522680,-94.79138540
2038,Lara,239,VE,Venezuela,K,state,33.98221650,-118.13227470
2053,Mérida,239,VE,Venezuela,L,state,20.96737020,-89.59258570
2037,Miranda,239,VE,Venezuela,M,state,42.35193830,-71.52907660
2054,Monagas,239,VE,Venezuela,N,state,9.32416520,-63.01475780
2052,"Nueva Esparta",239,VE,Venezuela,O,state,10.99707230,-63.91132960
2036,Portuguesa,239,VE,Venezuela,P,state,9.09439990,-69.09702300
2056,Sucre,239,VE,Venezuela,R,state,-19.03534500,-65.25921280
2048,Táchira,239,VE,Venezuela,S,state,7.91370010,-72.14161320
2043,Trujillo,239,VE,Venezuela,T,state,36.67343430,-121.62875880
2041,Yaracuy,239,VE,Venezuela,U,state,10.33938900,-68.81088490
2042,Zulia,239,VE,Venezuela,V,state,10.29102370,-72.14161320
3794,"An Giang",240,VN,Vietnam,44,,10.52158360,105.12589550
3770,"Bà Rịa-Vũng Tàu",240,VN,Vietnam,43,,10.54173970,107.24299760
3815,"Bắc Giang",240,VN,Vietnam,54,,21.28199210,106.19747690
3822,"Bắc Kạn",240,VN,Vietnam,53,,22.30329230,105.87600400
3804,"Bạc Liêu",240,VN,Vietnam,55,,9.29400270,105.72156630
3791,"Bắc Ninh",240,VN,Vietnam,56,,21.12144400,106.11105010
3796,"Bến Tre",240,VN,Vietnam,50,,10.24335560,106.37555100
3785,"Bình Dương",240,VN,Vietnam,57,,11.32540240,106.47701700
3830,"Bình Định",240,VN,Vietnam,31,,14.16653240,108.90268300
3797,"Bình Phước",240,VN,Vietnam,58,,11.75118940,106.72346390
3787,"Bình Thuận",240,VN,Vietnam,40,,11.09037030,108.07207810
3778,"Cà Mau",240,VN,Vietnam,59,,9.15267280,105.19607950
4925,"Cần Thơ",240,VN,Vietnam,CT,,10.03418510,105.72255070
3782,"Cao Bằng",240,VN,Vietnam,04,,22.63568900,106.25221430
3806,"Đà Nẵng",240,VN,Vietnam,DN,,16.05440680,108.20216670
3829,"Đắk Lắk",240,VN,Vietnam,33,,12.71001160,108.23775190
3823,"Đắk Nông",240,VN,Vietnam,72,,12.26464760,107.60980600
3773,"Điện Biên",240,VN,Vietnam,71,,21.80423090,103.10765250
3821,"Đồng Nai",240,VN,Vietnam,39,,11.06863050,107.16759760
3769,"Đồng Tháp",240,VN,Vietnam,45,,10.49379890,105.68817880
3813,"Gia Lai",240,VN,Vietnam,30,,13.80789430,108.10937500
3779,"Hà Giang",240,VN,Vietnam,03,,22.80255880,104.97844940
3802,"Hà Nam",240,VN,Vietnam,63,,20.58351960,105.92299000
3810,"Hà Nội",240,VN,Vietnam,HN,,21.02776440,105.83415980
3816,"Hà Tĩnh",240,VN,Vietnam,23,,18.35595370,105.88774940
3827,"Hải Dương",240,VN,Vietnam,61,,20.93734130,106.31455420
3783,"Hải Phòng",240,VN,Vietnam,HP,,20.84491150,106.68808410
3777,"Hậu Giang",240,VN,Vietnam,73,,9.75789800,105.64125270
3811,"Hồ Chí Minh",240,VN,Vietnam,SG,,10.82309890,106.62966380
3799,"Hòa Bình",240,VN,Vietnam,14,,20.68612650,105.31311850
3768,"Hưng Yên",240,VN,Vietnam,66,,20.85257110,106.01699710
3793,"Khánh Hòa",240,VN,Vietnam,34,,12.25850980,109.05260760
3800,"Kiên Giang",240,VN,Vietnam,47,,9.82495870,105.12589550
3772,"Kon Tum",240,VN,Vietnam,28,,14.34974030,108.00046060
3825,"Lai Châu",240,VN,Vietnam,01,,22.38622270,103.47026310
3818,"Lâm Đồng",240,VN,Vietnam,35,,11.57527910,108.14286690
3792,"Lạng Sơn",240,VN,Vietnam,09,,21.85370800,106.76151900
3817,"Lào Cai",240,VN,Vietnam,02,,22.48094310,103.97549590
3808,"Long An",240,VN,Vietnam,41,,10.56071680,106.64976230
3789,"Nam Định",240,VN,Vietnam,67,,20.43882250,106.16210530
3780,"Nghệ An",240,VN,Vietnam,22,,19.23424890,104.92003650
3786,"Ninh Bình",240,VN,Vietnam,18,,20.25061490,105.97445360
3788,"Ninh Thuận",240,VN,Vietnam,36,,11.67387670,108.86295720
3801,"Phú Thọ",240,VN,Vietnam,68,,21.26844300,105.20455730
3824,"Phú Yên",240,VN,Vietnam,32,,13.08818610,109.09287640
3809,"Quảng Bình",240,VN,Vietnam,24,,17.61027150,106.34874740
3776,"Quảng Nam",240,VN,Vietnam,27,,15.53935380,108.01910200
3828,"Quảng Ngãi",240,VN,Vietnam,29,,15.12138730,108.80441450
3814,"Quảng Ninh",240,VN,Vietnam,13,,21.00638200,107.29251440
3803,"Quảng Trị",240,VN,Vietnam,25,,16.74030740,107.18546790
3819,"Sóc Trăng",240,VN,Vietnam,52,,9.60252100,105.97390490
3812,"Sơn La",240,VN,Vietnam,05,,21.10222840,103.72891670
3826,"Tây Ninh",240,VN,Vietnam,37,,11.33515540,106.10988540
3775,"Thái Bình",240,VN,Vietnam,20,,20.44634710,106.33658280
3807,"Thái Nguyên",240,VN,Vietnam,69,,21.56715590,105.82520380
3771,"Thanh Hóa",240,VN,Vietnam,21,,19.80669200,105.78518160
3798,"Thừa Thiên-Huế",240,VN,Vietnam,26,,16.46739700,107.59053260
3781,"Tiền Giang",240,VN,Vietnam,46,,10.44933240,106.34205040
3805,"Trà Vinh",240,VN,Vietnam,51,,9.81274100,106.29929120
3795,"Tuyên Quang",240,VN,Vietnam,07,,21.77672460,105.22801960
3790,"Vĩnh Long",240,VN,Vietnam,49,,10.23957400,105.95719280
3774,"Vĩnh Phúc",240,VN,Vietnam,70,,21.36088050,105.54743730
3784,"Yên Bái",240,VN,Vietnam,06,,21.71676890,104.89858780
5074,"Saint Croix",242,VI,"Virgin Islands (US)",SC,district,17.72935200,-64.73437050
5073,"Saint John",242,VI,"Virgin Islands (US)",SJ,district,18.33561690,-64.80028000
5072,"Saint Thomas",242,VI,"Virgin Islands (US)",ST,district,18.34284590,-65.07701800
1242,'Adan,245,YE,Yemen,AD,,12.82574810,44.79438040
1250,'Amran,245,YE,Yemen,AM,,16.25692140,43.94367880
1237,Abyan,245,YE,Yemen,AB,,13.63434130,46.05632120
1240,"Al Bayda'",245,YE,Yemen,BA,,14.35886620,45.44980650
1241,"Al Hudaydah",245,YE,Yemen,HU,,15.30530720,43.01948970
1243,"Al Jawf",245,YE,Yemen,JA,,16.79018190,45.29938620
1251,"Al Mahrah",245,YE,Yemen,MR,,16.52384230,51.68342750
1235,"Al Mahwit",245,YE,Yemen,MW,,15.39632290,43.56069460
1232,"Amanat Al Asimah",245,YE,Yemen,SA,,15.36944510,44.19100660
1246,Dhamar,245,YE,Yemen,DH,,14.71953440,44.24790150
1238,Hadhramaut,245,YE,Yemen,HD,,16.93041350,49.36531490
1244,Hajjah,245,YE,Yemen,HJ,,16.11806310,43.32946600
1233,Ibb,245,YE,Yemen,IB,,14.14157170,44.24790150
1245,Lahij,245,YE,Yemen,LA,,13.14895880,44.85054950
1234,Ma'rib,245,YE,Yemen,MA,,15.51588800,45.44980650
1248,Raymah,245,YE,Yemen,RA,,14.62776820,43.71424840
1249,Saada,245,YE,Yemen,SD,,16.84765280,43.94367880
1236,Sana'a,245,YE,Yemen,SN,,15.31689130,44.47480180
1247,Shabwah,245,YE,Yemen,SH,,14.75463030,46.51626200
1239,Socotra,245,YE,Yemen,SU,,12.46342050,53.82373850
1231,Ta'izz,245,YE,Yemen,TA,,13.57758860,44.01779890
1986,"Central Province",246,ZM,Zambia,02,,7.25649960,80.72144170
1984,"Copperbelt Province",246,ZM,Zambia,08,,-13.05700730,27.54958460
1991,"Eastern Province",246,ZM,Zambia,03,,23.16696880,49.36531490
1987,"Luapula Province",246,ZM,Zambia,04,,-11.56483100,29.04599270
1988,"Lusaka Province",246,ZM,Zambia,09,,-15.36571290,29.23207840
1989,"Muchinga Province",246,ZM,Zambia,10,,-15.38219300,28.26158000
1982,"Northern Province",246,ZM,Zambia,05,,8.88550270,80.27673270
1985,"Northwestern Province",246,ZM,Zambia,06,,-13.00502580,24.90422080
1990,"Southern Province",246,ZM,Zambia,07,,6.23737500,80.54384500
1983,"Western Province",246,ZM,Zambia,01,,6.90160860,80.00877460
1956,"Bulawayo Province",247,ZW,Zimbabwe,BU,,-20.14895050,28.53310380
1958,"Harare Province",247,ZW,Zimbabwe,HA,,-17.82162880,31.04922590
1959,Manicaland,247,ZW,Zimbabwe,MA,,-18.92163860,32.17460500
1955,"Mashonaland Central Province",247,ZW,Zimbabwe,MC,,-16.76442950,31.07937050
1951,"Mashonaland East Province",247,ZW,Zimbabwe,ME,,-18.58716420,31.26263660
1953,"Mashonaland West Province",247,ZW,Zimbabwe,MW,,-17.48510290,29.78892480
1960,"Masvingo Province",247,ZW,Zimbabwe,MV,,-20.62415090,31.26263660
1954,"Matabeleland North Province",247,ZW,Zimbabwe,MN,,-18.53315660,27.54958460
1952,"Matabeleland South Province",247,ZW,Zimbabwe,MS,,-21.05233700,29.04599270
1957,"Midlands Province",247,ZW,Zimbabwe,MI,,-19.05520090,29.60354950
`

func (store *Store) seedStatesIfTableEmpty() error {
	list, err := store.StateList(StateQueryOptions{
		Limit: 1,
	})

	if err != nil {
		return err
	}

	if len(list) > 0 {
		return nil
	}

	// Create a new CSV reader from the string data
	reader := csv.NewReader(strings.NewReader(strings.TrimSpace(statesCsvData)))

	// Read all the records from the CSV reader
	states, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, entry := range states {
		// id,name,country_id,country_code,country_name,state_code,type,latitude,longitude
		name := entry[1]
		countryCode := entry[3]
		stateCode := entry[5]

		state := NewState()
		state.SetCountryCode(countryCode)
		state.SetStateCode(stateCode)
		state.SetName(name)

		err = store.StateCreate(state)

		if err != nil {
			return err
		}
	}

	return nil
}
