///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package utils

import "gitlab.com/xx_network/primitives/id"

var ServerID = id.ID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var Contract = `<style>
    div.agreement p, div.agreement a, div.agreement strong, div.agreement ol, div.agreement li {
        margin: 0;
        padding: 0;
        border: 0;
        vertical-align: baseline;
    }

    div.agreement {
        line-height: 1;
        font-family: TimesNewRoman, "Times New Roman", Times, Baskerville, Georgia, serif;
        padding: 2em;
    }

    div.agreement p {
        margin: 1em 0;
        line-height: 1.25em;
        text-align: justify;
    }

    div.agreement strong {
        font-weight: 700;
    }

    div.agreement ol {
        padding-left: 40px
    }

    div.agreement ol > li > ol > li {
        padding-left: 10px;
    }

    div.agreement ol[type="1"] > li p {
        margin: 0.5em 0;
    }

    div.agreement ol[type="1"] {
        list-style-type: none;
        counter-reset: item;
        margin: 0;
        padding: 0;
    }

    div.agreement ol[type="1"] > li {
        display: table;
        counter-increment: item;
        margin-bottom: 0.5em;
    }

    div.agreement ol[type="1"] > li:before {
        content: counters(item, ".") ". ";
        display: table-cell;
        padding-right: 10px;
        text-align: right;
        width: 1.25em;
    }

    div.agreement li ol[type="1"] > li {
        margin: 0;
    }

    div.agreement li ol[type="1"] > li:before {
        content: counters(item, ".") ". ";
        text-align: right;
    }
</style>
<div class="agreement">
	<p style="text-align: center;"><strong>PARTICIPANT TERMS AND CONDITIONS FOR MAINNET TRANSITION PROGRAM</strong></p>
	<p><strong>PLEASE READ THE FOLLOWING TERMS AND CONDITIONS CAREFULLY BEFORE DOWNLOADING, INSTALLING, OR USING ANY NODE SOFTWARE OR OTHERWISE PARTICIPATING IN THE MAINNET TRANSITION PROGRAM.</strong></p>
	<p><strong>BY CLICKING TO SIGNIFY ACCEPTANCE OR OTHERWISE DOWNLOADING, INSTALLING, RUNNING OR USING THE XX NETWORK OR NODE SOFTWARE OR OTHERWISE PARTICIPATING IN THE MAINNET TRANSITION PROGRAM, YOU AGREE TO BE BOUND BY THE TERMS OF THIS AGREEMENT EFFECTIVE AS OF THE DATE THAT YOU TAKE THE EARLIEST OF ONE OF THE FOREGOING ACTIONS (“Effective Date</strong>”)<strong>.</strong></p>
	<p><strong>About this Agreement</strong>. These Participant Terms and Conditions for MainNet Transition Program (this “<strong>Agreement</strong>”) constitute a legally binding agreement between you (the “<strong>Participant</strong>” or “<strong>You</strong>” or derivatives thereof) and xx Labs SEZC, an entity incorporated under the laws of the Cayman Islands (“<strong>Company</strong>”) and governs your participation in the MainNet Transition Program. This Agreement supersedes and merges all prior oral and written agreements, discussions and understandings between you and the Company with respect to its subject matter, and neither of the parties will be bound by any conditions, inducements or representations other than as expressly provided for herein.</p>
	<p><strong>About XX Network</strong>. The “<strong>xx network</strong>” is a quantum ready blockchain network designed to enable a scalable and private payments and messaging system. The xx network is a nominated proof-of-stake blockchain, with Participants and their nominators compensated with Coins for producing blocks and running the protocol.</p>
	<p><strong>About MainNet Transition Program</strong>. The “<strong>MainNet Transition Program</strong>” is described at <a href="https://xx.network/mainnet-transition" rel="noopener" target="_blank">https://xx.network/mainnet-transition</a>.</p>
	<p><strong>Participants</strong>. You, as the Participant, agree to operate one or more MainNet Transition Program Participant’s Nodes in the xx network and to be legally bound by this Agreement. You represent and warrant that you are at least 18 years old (or the age of majority as determined by the laws of your place of residency.</p>
	<p><strong>Electronic Signature and Disclosure Consent Notice</strong>. You agree to the use of electronic documents and records in connection with participation in the xx network—including without limitation this electronic signature and disclosure notice—and that this use satisfies any requirement to provide you these documents and their content in writing. <strong>If You do not agree, do not accept this Agreement</strong>.</p>
	<p><strong>Changes to this Agreement</strong>. The Company reserves the right to change this Agreement from time to time in our sole discretion. If the Company makes material changes to this Agreement, we will provide notice of such changes, such as by posting the revised Participant Terms and Conditions. By continuing to access or use the Node Software or otherwise participate in the xx network after the posted effective date of modifications to this Agreement, you agree to be bound by the revised version of the Agreement. If you do not agree to the modified Agreement, you must cease your participation the xx network and the Company reserves the right to terminate your participation.</p>
	<ol type="1">
		<li>
			<p><strong>Definitions</strong>. Capitalized terms not otherwise defined in the body of this Agreement have the following meanings:</p>
			<ol type="1">
				<li><p>“<strong>Coin</strong>” means a digital coin (i.e., token) created by Company as a cryptographically secured representation of the right to exchange or use such coin or token as payment for applications or services.</p></li>
				<li><p>“<strong>MainNet Launch Date</strong>” means the public launch of the xx network as determined by Company.</p></li>
				<li><p>“<strong>MainNet Transition Program Participant’s Node</strong>” or “<strong>Node</strong>” means an active physical electronic device that runs the Node Software on the xx network mainnet.</p></li>
				<li><p>“<strong>Node Hardware Requirements</strong>” means the minimum MainNet Transition Program Participant’s Node hardware and operating requirements posted at <a href="https://xx.network/mainnet-transition#mainnet-hardware-requirements" rel="noopener" target="_blank">https://xx.network/mainnet-transition#mainnet-hardware-requirements</a> and <a href="https://xxnetwork.wiki/index.php/Hardware_Requirements" rel="noopener" target="_blank">https://xxnetwork.wiki/index.php/Hardware_Requirements</a>, as such requirements may be changed by the Company from time to time, all of which are incorporated and made a part of this Agreement by this reference.</p></li>
				<li><p>“<strong>Node Rules</strong>” means the rules for participation in the xx network and use of the Coins, as such rules are posted on the xx network site, including at <a href="https://xx.network/mainnet-transition" rel="noopener" target="_blank">https://xx.network/mainnet-transition</a>, from time to time by the Company, all of which are incorporated and made a part of this Agreement by this reference.</p></li>
				<li><p>“<strong>Node Software</strong>” means the software provided or otherwise made available by Company that creates, sends, receives, or transmits information to the xx network, runs the consensus algorithm of the xx network, and offers one or more applications or services supported by xx network.</p></li>
			</ol>
		</li>
		<li><p><strong>Term</strong>. The term of this Agreement commences upon the Effective Date and unless terminated earlier in accordance with this Agreement continues for three (3) years after the MainNet Launch Date unless extended by the Company upon notice to the Participant.</p></li>
		<li>
			<p><strong>MainNet Participant Requirements</strong>.</p>
			<ol type="1">
				<li><p><span style="text-decoration: underline">Node Set-up.</span> The Participant will be required to set-up and manage its MainNet Transition Program Participant’s Node independently no later than three (3) months after MainNet Launch Date.</p></li>
				<li><p><span style="text-decoration: underline">Length of Operation.</span> The Participant may discontinue its participation in the MainNet Transition Program at any time in accordance with Section 10 (Discontinuation; Termination), in which case this Agreement terminates.</p></li>
				<li><p><span style="text-decoration: underline">Node Software and Operating Requirements.</span> MainNet Transition Program Participant’s Nodes must download and run the Node Software. The Participant will operate its MainNet Transition Program Participant’s Node, and use the Node Software and xx network, in accordance with the Node Rules. The Participant’s MainNet Transition Program Participant’s Node must meet the Node Hardware Requirements at all times. The Participant will keep its MainNet Transition Program Participant’s Node running with both cMix and xx consensus protocols. The Participant must use the provided software and system images as provided by the Company (e.g., Participant is prohibited from installing a different version of Python).</p></li>
				<li><p><span style="text-decoration: underline">Node Accounts.</span> The Participant will own and control its own MainNet Transition Program Participant’s Node account, including any Coins within such MainNet Transition Program Participant’s Node account.</p></li>
			</ol>
		</li>
		<li>
			<p><strong>Prohibited Activities; Rules of Conduct</strong>. As a condition of operating a MainNet Transition Program Participant’s Node, the Participant will not operate a MainNet Transition Program Participant’s Node for any purpose that is prohibited by this Agreement or the Node Rules. The Participant is responsible for all of its activity with regard to the acts and omissions of its MainNet Transition Program Participant’s Node or its representatives. The Participant shall not (directly or indirectly):</p>
			<ol type="a">
				<li><p>take any action that imposes or may impose (as determined by the Company in its sole discretion) an unreasonable or disproportionately large load on the Company’s (or its third-party providers’) infrastructure;</p></li>
				<li><p>interfere or attempt to interfere with the proper working of the xx network or any activities conducted on the xx network;</p></li>
				<li><p>run any software on the MainNet Transition Program Participant’s Node that is not approved by the Company; provided, however, the Participant may modify the Node Software subject to the Company’s right to object to any fork or other modification of the Node Software;</p></li>
				<li><p>use the xx network in a manner that is knowingly unlawful or fraudulent;</p></li>
				<li><p>share with any person or entity any non-public information provided by or on behalf of the Company to the Participant, including but not limited to the Node Software and/or related documentation or instructions or any generated cryptographic private keys;</p></li>
				<li><p>replace the public keys, Transport Layer Security (TLS) certifications, cyclic groups, or addresses provided by the Company;</p></li>
				<li><p>take any actions to create any claim, encumbrance or lien with respect to the xx network or the Node Software; or</p></li>
				<li><p>otherwise take any action in violation of the Node Rules.</p></li>
			</ol>
		</li>
		<li><p><strong>Use Rights</strong>. Subject to all of the terms and conditions of this Agreement, the Participant has a limited, non-exclusive, non-sublicensable and non-transferrable license to use the xx network for the Participant’s own personal use in connection with the MainNet Transition Program. The Node Software consists of open source code and is made available to the Participant pursuant to the open source license agreements that accompany or are otherwise available with the Node Software. The Participants use of the Node Software and participation in the MainNet Transition Program is subject to the Participant’s compliance with such open source license agreements. If Participant modifies or creates forks for the Node Software, the Company reserves the right to object to use of any such modifications or forks in connection with the xx network or otherwise. Use of the xx network and/or the Node Software and related documentation may be protected by applicable copyright, patent, and trademark Laws, international treaties, and other applicable laws. The Participant must not alter or remove any proprietary markings, including copyright, patent, trademark, service mark, trade secret, confidentiality or other proprietary notices. All rights not expressly granted to the Participant are reserved, including all worldwide rights in or to all patents, rights associated with works of authorship (including copyrights and similar rights), trade secrets and other proprietary rights arising under statutory or common law.</p></li>
		<li><p><strong>Stake Removal</strong>. If the Participant (or its MainNet Transition Program Participant’s Node) fails to comply with this Agreement, whether due to multiple failures of such MainNet Transition Program Participant’s Node, the severity of any one or multiple failures of such MainNet Transition Program Participant’s Node, or the duration of any such failure(s) (each, a “<strong>Node Failure</strong>”) and the Company believes in its reasonable sole discretion, that such failure(s) may cause a negative effect on the xx network or may subject the Company or any Participants, application developers or end users to any legal, compliance or business risks, the Company has the right to (i) remove or reduce the stake provided by the Company, if any, and/or (ii) terminate this Agreement in full upon written notice to the Participant.</p></li>
		<li><p><strong>Compensation</strong>. The Node Rules describe how the Participant may earn Coins. The Participant is solely responsible for any and all taxes, duties or similar payments that may be due related to the Coins earned or your participation in the xx network. YOU ACKNOWLEDGE AND AGREE THAT (i) THE VALUE OF COINS MAY CHANGE OR HAVE NO VALUE, (ii) THERE MAY BE NO MECHANISM OR MARKET FOR THE USE OR EXCHANGE OF THE COINS, and (iii) THE RATE OF NOMINATION ON A NODE FROM OTHER PARTIES MAY TAKE A SIGNIFICANT PORTION OF THE VALUE OF THE NODE’S EARNED COINS.</p></li>
		<li><p><strong>DISCLAIMER</strong>. THE NODE SOFTWARE, THE MAINNET TRANSITION PROGRAM AND XX NETWORK, INCLUDING WITHOUT LIMITATION THIRD-PARTY CODE, ARE PROVIDED “AS IS” AND “AS AVAILABLE.” COMPANY DISCLAIMS ALL REPRESENTATIONS AND WARRANTIES, EXPRESS OR IMPLIED, STATUTORY OR OTHERWISE, INCLUDING BUT NOT LIMITED TO THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, TITLE, NON-INFRINGEMENT, ACCURACY, SATISFACTORY QUALITY, AND QUIET ENJOYMENT. COMPANY MAKES NO WARRANTY THAT THE NODE SOFTWARE, THE MAINNET TRANSITION PROGRAM OR XX NETWORK WILL BE UNINTERRUPTED, ACCURATE, COMPLETE, RELIABLE, CURRENT, ERROR-FREE, VIRUS FREE, OR FREE OF MALICIOUS CODE OR HARMFUL COMPONENTS, OR THAT DEFECTS WILL BE CORRECTED. COMPANY DOES NOT CONTROL, ENDORSE, SPONSOR, OR ADOPT ANY MATERIALS, CONTENT, MESSAGES OR COMMUNICATIONS AND MAKES NO REPRESENTATIONS OR WARRANTIES OF ANY KIND REGARDING THE SAME. COMPANY HAS NO OBLIGATION TO SCREEN, MONITOR, OR EDIT AND IS NOT RESPONSIBLE OR LIABLE FOR ANY MATERIALS, CONTENT, MESSAGES OR COMMUNICATIONS. YOU ACKNOWLEDGE AND AGREE THAT COMPANY HAS NO INDEMNITY, SUPPORT, SERVICE LEVEL, OR OTHER OBLIGATIONS HEREUNDER.</p></li>
		<li><p><strong>Limitations on Liability</strong>. NEITHER THE COMPANY NOR ANY REPRESENTATIVE OF THE COMPANY SHALL BE LIABLE FOR ANY CONSEQUENTIAL, INDIRECT, SPECIAL, PUNITIVE, INCIDENTAL OR EXEMPLARY DAMAGES, OR DAMAGES FOR LOST DATA, SOFTWARE, FIRMWARE, LOST PROFITS, BUSINESS OR REVENUES, WHETHER FORESEEABLE OR UNFORESEEABLE, EVEN IF SUCH PARTY HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGES, AND IN NO EVENT WILL COMPANY’S TOTAL AGGREGATE LIABILITY ARISING FROM OR RELATING TO THIS AGREEMENT EXCEED US $25.</p></li>
		<li><p><strong>Discontinuation; Termination</strong>. The Company reserves the right to discontinue the MainNet Transition Program and/or terminate this Agreement and/or Participant’s participation in the MainNet Transition Program at any time, for any or no reason, in its sole discretion by providing written notice at least 20 days in advance to the Participant. The Participant may terminate this Agreement at any time, for any or no reason, in its sole discretion by providing written notice at least 20 days in advance to the Company. Subject to Section 6 (Stake Removal), if participation in the MainNet Transition Program is discontinued or the Agreement terminated for any reason, the Participant may retain its Coins and the foregoing represents the Participant’s sole remedy.</p></li>
		<li><p><strong>Indemnity</strong>. To the fullest extent permitted by applicable law, the Participant will defend (at the Company’s request), indemnify and hold harmless Company, its affiliates and their respective past, present, and future employees, officers, directors, contractors, consultants, equity holders, suppliers, vendors, service providers, parent companies, subsidiaries, affiliates, agents, representatives, predecessors, successors and assigns from and against all claims, damages, costs, liabilities and expenses (including attorneys’ fees) that arise from or relate to: (i) your use of the Node Software; (ii) your participation in the MainNet Transition Program; (iii) any materials, content, messaging or communications you provide; (iv) any Node Failure; or (v) your breach of any of this Agreement.</p></li>
		<li><p><strong>Notices</strong>. Any notice required or permitted by this Agreement will be deemed sufficient when delivered to (i) you via email to the email address you provided during registration and (ii) the Company when delivered by overnight courier to
			<br />
			Hermes Corporate Service Ltd., Fifth Floor
			<br />
			Zephyr House, P.O. Box 31493, George Town
			<br />
			Grand Cayman KY1-1206, attention, CEO, as such may be subsequently modified by notice.</p></li>
		<li><p><strong>Governing Law</strong>. This Agreement and all rights and obligations hereunder will be governed by the laws of the Cayman Islands, without regard to the conflicts of law provisions of such jurisdiction. The application of the United Nations Convention on Contracts for the International Sale of Goods is expressly excluded.</p></li>
		<li><p><strong>Assignment</strong>. This Agreement is not transferable or assignable, by operation of law or otherwise, by the Participant. The Company reserves the right to assign, without the Participant’s consent, this Agreement by operation of law or otherwise to (i) an affiliate or (ii) a third party in connection with the sale, change or control, reorganization, merger or other business combination of the business to which this Agreement relates. Subject to the foregoing, the Agreement shall be binding upon the parties and their respective administrators, successors and assigns.</p></li>
		<li><p><strong>Severability</strong>. If any one or more of the provisions of this Agreement is for any reason held to be invalid, illegal or unenforceable, in whole or in part or in any respect, or if any one or more of the provisions of this Agreement operate or would prospectively operate to invalidate this Agreement, then and in any such event, such provision(s) only will be deemed null and void and will not affect any other provision of this Agreement and the remaining provisions of this Agreement will remain operative and in full force and effect and will not be affected, prejudiced, or disturbed thereby.</p></li>
		<li><p><strong>No Partnership; Third Party Beneficiaries</strong>. This Agreement does not create a partnership, agency relationship, or joint venture between the parties. Participant’s relationship with the Company will be that of an independent contractor and not that of an employee. This Agreement is intended solely for the benefit of the parties and are not intended to confer third-party beneficiary rights upon any other person or entity other than those persons or entities afforded the protections of Section 11 (Indemnity).</p></li>
		<li><p><strong>Waiver</strong>. Failure of the Company to enforce a right under this Agreement shall not act as a waiver of that right or the ability to later assert that right relative to the particular situation involved.</p></li>
		<li><p><strong>Export Control Laws</strong>. The Participant acknowledges that the use of the xx network, the Node Software and all related technical information, documents, and materials may be subject to import and export regulations and controls, as well as end user, end use, and destination restrictions issued by governments. The Participant must comply strictly with all export controls and shall not export, re-export, transfer or divert any of the Node Software or any direct product thereof, to any destination, end use, or end user that is prohibited or restricted under any export control laws and regulations.</p></li>
	</ol>
</div>`

const NovemberContract = `<style>
    div.agreement p, div.agreement a, div.agreement strong, div.agreement ol, div.agreement li {
        margin: 0;
        padding: 0;
        border: 0;
        vertical-align: baseline;
    }

    div.agreement {
        line-height: 1;
        font-family: TimesNewRoman, "Times New Roman", Times, Baskerville, Georgia, serif;
        padding: 2em;
    }

    div.agreement p {
        margin: 1em 0;
        line-height: 1.25em;
        text-align: justify;
    }

    div.agreement strong {
        font-weight: 700;
    }

    div.agreement ol {
        padding-left: 40px
    }

    div.agreement ol > li > ol > li {
        padding-left: 10px;
    }

    div.agreement ol[type="1"] > li p {
        margin: 0.5em 0;
    }

    div.agreement ol[type="1"] {
        list-style-type: none;
        counter-reset: item;
        margin: 0;
        padding: 0;
    }

    div.agreement ol[type="1"] > li {
        display: table;
        counter-increment: item;
        margin-bottom: 0.5em;
    }

    div.agreement ol[type="1"] > li:before {
        content: counters(item, ".") ". ";
        display: table-cell;
        padding-right: 10px;
        text-align: right;
        width: 1.25em;
    }

    div.agreement li ol[type="1"] > li {
        margin: 0;
    }

    div.agreement li ol[type="1"] > li:before {
        content: counters(item, ".") ". ";
        text-align: right;
    }
</style>
<div class="agreement">
	<p style="text-align: center"><strong>TERMS AND CONDITIONS FOR MAINNET SUPPORT REIMBURSEMENT</strong></p>
	<p><strong>PLEASE READ THE FOLLOWING TERMS AND CONDITIONS CAREFULLY (“TERMS”). BY ACCEPTING THE PROPOSED CONSIDERATION YOU AGREE TO BE BOUND BY THESE TERMS.</strong></p>
	<p><strong>About this Agreement</strong>. These Terms and Conditions for MainNet Support Reimbursement (this “<strong>Agreement</strong>”) constitute a legally binding agreement between you (the “<strong>Node Operator</strong>” or “<strong>You</strong>” or derivatives thereof) and xx Labs SEZC, an entity incorporated under the laws of the Cayman Islands (“<strong>Company</strong>”) and governs your relationship with the Company. This Agreement supersedes and merges all prior oral and written agreements, discussions and understandings between you and the Company with respect to its subject matter, and neither of the parties will be bound by any conditions, inducements or representations other than as expressly provided for herein.</p>
	<p><strong>About XX Network</strong>. The “<strong>xx network</strong>” is a quantum ready blockchain network designed to enable a scalable and private payments and messaging system. The xx network is a nominated proof-of-stake blockchain, with Node Operator and their nominators compensated with Coins for producing blocks and running the protocol.</p>
	<p><strong>Node Operator</strong>. You, as the Node Operator, agree to operate one or more MainNet Nodes in the xx network and to be legally bound by this Agreement. You represent and warrant that you are at least 18 years old (or the age of majority as determined by the laws of your place of residency.</p>
	<p><strong>Electronic Signature and Disclosure Consent Notice</strong>. You agree to the use of electronic documents and records in connection with participation in the xx network—including without limitation this electronic signature and disclosure notice—and that this use satisfies any requirement to provide you these documents and their content in writing. <strong>If You do not agree, do not accept this Agreement</strong>.</p>
	<p><strong>Changes to this Agreement</strong>. The Company reserves the right to change this Agreement from time to time in our sole discretion. If the Company makes material changes to this Agreement, we will provide notice of such changes, such as by posting the revised Node Operator Terms and Conditions. By continuing to act as a Node Operator and hosting a MainNet Node, after the posted effective date of modifications to this Agreement, you agree to be bound by the revised version of the Agreement. If you do not agree to the modified Agreement, you must cease your hosting of a MainNet Node</p>
	<ol type="1">
		<li>
			<p><strong>Definitions</strong>. Capitalized terms not otherwise defined in the body of this Agreement have the following meanings:</p>
			<ol type="1">
				<li><p>“<strong>Coin</strong>” means a digital coin (i.e., token) created by Company as a cryptographically secured representation of the right to exchange or use such coin or token as payment for applications or services.</p></li>
				<li><p>“<strong>MainNet Launch Date</strong>” means November 17, 2021.</p></li>
				<li><p>“<strong>MainNet Node</strong>” or “<strong>Node</strong>” means an active physical electronic device that runs the Node Software on the xx network mainnet.</p></li>
				<li><p>“<strong>Node Hardware Requirements</strong>” means the minimum MainNet Node hardware and operating requirements posted at <a href="https://xx.network/mainnet-transition#mainnet-hardware-requirements" rel="noopener" target="_blank">https://xx.network/mainnet-transition#mainnet-hardware-requirements</a> and <a href="https://xxnetwork.wiki/index.php/Hardware_Requirements" rel="noopener" target="_blank">https://xxnetwork.wiki/index.php/Hardware_Requirements</a>, as such requirements may be changed by the Company from time to time, all of which are incorporated and made a part of this Agreement by this reference.</p></li>
				<li><p>“<strong>Node Rules</strong>” means the rules for Node Operator in the xx network and use of the Coins, as such rules are posted on the xx network site, including at <a href="https://xx.network/mainnet-transition" rel="noopener" target="_blank">https://xx.network/mainnet-transition</a>, from time to time by the Company, all of which are incorporated and made a part of this Agreement by this reference.</p></li>
				<li><p>“<strong>Node Software</strong>” means the software provided or otherwise made available by Company that creates, sends, receives, or transmits information to the xx network, runs the consensus algorithm of the xx network, and offers one or more applications or services supported by xx network.</p></li>
			</ol>
		</li>
		<li><p><strong>Reimbursement</strong>. The Company recognizes that Node Operators that fulfilled the MainNet Node Operator Requirements below leading up to the MainNet Launch Date and through the month of November 2021, provided benefits to the Company and MainNet by maintaining and operating a MainNet Node. Accordingly, the Company seeks to compensate and reimburse such Node Operators the amount of 7,000 xx Coins each (the “Reimbursement Coins”). In connection with execution of this Agreement, Node Operators shall provide the Company with wallet information pursuant to the instructions provided by the Company. Within fifteen (15) business days of executing this Agreement and provided the requested wallet information, the Company will collect the Node Operator’s wallet and issue the Reimbursement Coins to the Node Operator’s wallet. Upon receipt, the Reimbursement Coins will be the property of Node Operator without restriction.</p></li>
		<li>
			<p><strong>MainNet Node Operator Requirements</strong>.</p>
			<ol type="1">
				<li><p><span style="text-decoration: underline">Node Set-up</span>. The Node Operator set-up and managed its MainNet Node independently before the MainNet Launch Date.</p></li>
				<li><p><span style="text-decoration: underline">Length of Operation</span>. The Node Operator may discontinue its support of MainNet at any time in accordance with Section 10 (Discontinuation; Termination), in which case this Agreement terminates.</p></li>
				<li><p><span style="text-decoration: underline">Node Software and Operating Requirements</span>. Node Operators must have downloaded and run the Node Software. The Node Operator has, and will continue to, operate its MainNet Node, and use the Node Software and xx network, in accordance with the Node Rules. The Node Operator’s Node must meet the Node Hardware Requirements at all times. The Node Operator will keep its MainNet Node running with both cMix and xx consensus protocols. The Node Operator must use the provided software and system images as provided by the Company (e.g., Node Operator is prohibited from installing a different version of Python).</p></li>
				<li><p><span style="text-decoration: underline">Node Accounts</span>. The Node Operator will own and control its own MainNet Node Operator account, including any Coins within such MainNet Node Operator account.</p></li>
			</ol>
		</li>
		<li>
			<p><strong>Prohibited Activities; Rules of Conduct</strong>. As a condition of operating a MainNet Node, the Node Operator will not operate a MainNet Node for any purpose that is prohibited by this Agreement or the Node Rules. The Node Operator shall not (directly or indirectly):</p>
			<ol type="a">
				<li><p>take any action that imposes or may impose (as determined by the Company in its sole discretion) an unreasonable or disproportionately large load on the Company’s (or its third-party providers’) infrastructure;</p></li>
				<li><p>interfere or attempt to interfere with the proper working of the xx network or any activities conducted on the xx network;</p></li>
				<li><p>run any software on the Node Operator’s Node that is not approved by the Company; provided, however, the Node Operator may modify the Node Software subject to the Company’s right to object to any fork or other modification of the Node Software;</p></li>
				<li><p>use the xx network in a manner that is knowingly unlawful or fraudulent;</p></li>
				<li><p>share with any person or entity any non-public information provided by or on behalf of the Company to the Node Operator, including but not limited to the Node Software and/or related documentation or instructions or any generated cryptographic private keys;</p></li>
				<li><p>replace the public keys, Transport Layer Security (TLS) certifications, cyclic groups, or addresses provided by the Company;</p></li>
				<li><p>take any actions to create any claim, encumbrance or lien with respect to the xx network or the Node Software; or</p></li>
				<li><p>otherwise take any action in violation of the Node Rules.</p></li>
			</ol>
		</li>
		<li><p><strong>Use Rights</strong>. Subject to all of the terms and conditions of this Agreement, the Node Operator has a limited, non-exclusive, non-sublicensable and non-transferrable license to use the xx network for the Node Operator’s own personal use in connection with MainNet. The Node Software consists of open source code and is made available to the Node Operator pursuant to the open source license agreements that accompany or are otherwise available with the Node Software. The Node Operator use of the Node Software and will comply with such open source license agreements covering the Node Software. If Node Operator modifies or creates forks for the Node Software, the Company reserves the right to object to use of any such modifications or forks in connection with the xx network or otherwise. Use of the xx network and/or the Node Software and related documentation may be protected by applicable copyright, patent, and trademark Laws, international treaties, and other applicable laws. The Node Operator must not alter or remove any proprietary markings, including copyright, patent, trademark, service mark, trade secret, confidentiality or other proprietary notices. All rights not expressly granted to the Node Operator are reserved, including all worldwide rights in or to all patents, rights associated with works of authorship (including copyrights and similar rights), trade secrets and other proprietary rights arising under statutory or common law.</p></li>
		<li><p><strong>Stake Removal</strong>. If the Node Operator fails to comply with this Agreement, the severity of any one or multiple failures, or the duration of any such failure(s) (each, a “<strong>Node Failure</strong>”) and the Company believes in its reasonable sole discretion, that such failure(s) may cause a negative effect on the xx network or may subject the Company or any Node Operators, application developers or end users to any legal, compliance or business risks, the Company has the right to (i) remove or reduce the stake provided by the Company, if any, and/or (ii) terminate this Agreement in full upon written notice to the Node Operator.</p></li>
		<li><p><strong>Compensation</strong>. Section 2 (Reimbursement) of this Agreement describes how Node Operator is being compensated for supporting MainNet in November 2021. The Node Operator is solely responsible for any and all taxes, duties or similar payments that may be due related to the Coins earned hereunder. YOU ACKNOWLEDGE AND AGREE THAT (i) THE VALUE OF COINS MAY CHANGE OR HAVE NO VALUE, (ii) THERE MAY BE NO MECHANISM OR MARKET FOR THE USE OR EXCHANGE OF THE COINS, and (iii) THE RATE OF NOMINATION ON A NODE FROM OTHER PARTIES MAY TAKE A SIGNIFICANT PORTION OF THE VALUE OF THE NODE’S EARNED COINS.</p></li>
		<li><p><strong>DISCLAIMER</strong>. THE NODE SOFTWARE AND XX NETWORK, INCLUDING WITHOUT LIMITATION THIRD-PARTY CODE, ARE PROVIDED “AS IS” AND “AS AVAILABLE.” COMPANY DISCLAIMS ALL REPRESENTATIONS AND WARRANTIES, EXPRESS OR IMPLIED, STATUTORY OR OTHERWISE, INCLUDING BUT NOT LIMITED TO THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, TITLE, NON-INFRINGEMENT, ACCURACY, SATISFACTORY QUALITY, AND QUIET ENJOYMENT. COMPANY MAKES NO WARRANTY THAT THE NODE SOFTWARE OR XX NETWORK WILL BE UNINTERRUPTED, ACCURATE, COMPLETE, RELIABLE, CURRENT, ERROR-FREE, VIRUS FREE, OR FREE OF MALICIOUS CODE OR HARMFUL COMPONENTS, OR THAT DEFECTS WILL BE CORRECTED. COMPANY DOES NOT CONTROL, ENDORSE, SPONSOR, OR ADOPT ANY MATERIALS, CONTENT, MESSAGES OR COMMUNICATIONS AND MAKES NO REPRESENTATIONS OR WARRANTIES OF ANY KIND REGARDING THE SAME. COMPANY HAS NO OBLIGATION TO SCREEN, MONITOR, OR EDIT AND IS NOT RESPONSIBLE OR LIABLE FOR ANY MATERIALS, CONTENT, MESSAGES OR COMMUNICATIONS. YOU ACKNOWLEDGE AND AGREE THAT COMPANY HAS NO INDEMNITY, SUPPORT, SERVICE LEVEL, OR OTHER OBLIGATIONS HEREUNDER.</p></li>
		<li><p><strong>Limitations on Liability</strong>. NEITHER THE COMPANY NOR ANY REPRESENTATIVE OF THE COMPANY SHALL BE LIABLE FOR ANY CONSEQUENTIAL, INDIRECT, SPECIAL, PUNITIVE, INCIDENTAL OR EXEMPLARY DAMAGES, OR DAMAGES FOR LOST DATA, SOFTWARE, FIRMWARE, LOST PROFITS, BUSINESS OR REVENUES, WHETHER FORESEEABLE OR UNFORESEEABLE, EVEN IF SUCH PARTY HAS BEEN ADVISED OF THE POSSIBILITY OF SUCH DAMAGES, AND IN NO EVENT WILL COMPANY’S TOTAL AGGREGATE LIABILITY ARISING FROM OR RELATING TO THIS AGREEMENT EXCEED US $25.</p></li>
		<li><p><strong>Discontinuation; Termination</strong>. The Company reserves the right to discontinue this Agreement and/or Node Operator’s participation as a MainNet Node Operator at any time, for any or no reason, in its sole discretion by providing written notice at least 20 days in advance to the Node Operator. Node Operator \ may terminate this Agreement at any time, for any or no reason, in its sole discretion by providing written notice at least 20 days in advance to the Company. Subject to Section 6 (Stake Removal), if Node Operator’s participation in MainNet as a Node Operator is discontinued or the Agreement terminated for any reason, the Node Operator may retain its Coins and the foregoing represents the Node Operator’s sole remedy.</p></li>
		<li><p><strong>Indemnity</strong>. To the fullest extent permitted by applicable law, the Node Operator will defend (at the Company’s request), indemnify and hold harmless Company, its affiliates and their respective past, present, and future employees, officers, directors, contractors, consultants, equity holders, suppliers, vendors, service providers, parent companies, subsidiaries, affiliates, agents, representatives, predecessors, successors and assigns from and against all claims, damages, costs, liabilities and expenses (including attorneys’ fees) that arise from or relate to: (i) your use of the Node Software; (ii) your participation as a Node Operator; (iii) any materials, content, messaging or communications you provide; (iv) any Node Failure; or (v) your breach of any of this Agreement.</p></li>
		<li><p><strong>Notices</strong>. Any notice required or permitted by this Agreement will be deemed sufficient when delivered to (i) you via email to the email address you provided during registration and (ii) the Company when delivered by overnight courier to Hermes Corporate Service Ltd., Fifth Floor, Zephyr House, P.O. Box 31493, George Town, Grand Cayman KY1-1206, attention, CEO, as such may be subsequently modified by notice.</p></li>
		<li><p><strong>Governing Law</strong>. This Agreement and all rights and obligations hereunder will be governed by the laws of the Cayman Islands, without regard to the conflicts of law provisions of such jurisdiction. The application of the United Nations Convention on Contracts for the International Sale of Goods is expressly excluded.</p></li>
		<li><p><strong>Assignment</strong>. This Agreement is not transferable or assignable, by operation of law or otherwise, by the Node Operator. The Company reserves the right to assign, without the Node Operator’s consent, this Agreement by operation of law or otherwise to (i) an affiliate or (ii) a third party in connection with the sale, change or control, reorganization, merger or other business combination of the business to which this Agreement relates. Subject to the foregoing, the Agreement shall be binding upon the parties and their respective administrators, successors and assigns.</p></li>
		<li><p><strong>Severability</strong>. If any one or more of the provisions of this Agreement is for any reason held to be invalid, illegal or unenforceable, in whole or in part or in any respect, or if any one or more of the provisions of this Agreement operate or would prospectively operate to invalidate this Agreement, then and in any such event, such provision(s) only will be deemed null and void and will not affect any other provision of this Agreement and the remaining provisions of this Agreement will remain operative and in full force and effect and will not be affected, prejudiced, or disturbed thereby.</p></li>
		<li><p><strong>No Partnership; Third Party Beneficiaries</strong>. This Agreement does not create a partnership, agency relationship, or joint venture between the parties. Node Operator’ relationship with the Company will be that of an independent contractor and not that of an employee. This Agreement is intended solely for the benefit of the parties and are not intended to confer third-party beneficiary rights upon any other person or entity other than those persons or entities afforded the protections of Section 11 (Indemnity).</p></li>
		<li><p><strong>Waiver</strong>. Failure of the Company to enforce a right under this Agreement shall not act as a waiver of that right or the ability to later assert that right relative to the particular situation involved.</p></li>
		<li><p><strong>Export Control Laws</strong>. The Node Operator acknowledges that the use of the xx network, the Node Software and all related technical information, documents, and materials may be subject to import and export regulations and controls, as well as end user, end use, and destination restrictions issued by governments. The Node Operator must comply strictly with all export controls and shall not export, re-export, transfer or divert any of the Node Software or any direct product thereof, to any destination, end use, or end user that is prohibited or restricted under any export control laws and regulations.</p></li>
	</ol>
</div>`
